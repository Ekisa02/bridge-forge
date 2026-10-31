import asyncio
from fastapi import FastAPI
from .api.v1 import router as api_router
from .config import settings
from .models.torch_model import TransactionClassifier
from .svc.web3_client import Web3Client


def create_app() -> FastAPI:
    app = FastAPI(title="Nebula Bridge", version="0.1.0")

    # include API router
    app.include_router(api_router, prefix="")

    # attach state objects
    @app.on_event("startup")
    async def startup():
        # load model (if available) or initialize randomly
        app.state.model = TransactionClassifier.load_from_checkpoint(settings.MODEL_PATH)
        app.state.web3 = Web3Client(provider_url=settings.WEB3_PROVIDER_URL)

        # start background telemetry polling
        app.state._telemetry_task = asyncio.create_task(_telemetry_poller(app))

    @app.on_event("shutdown")
    async def shutdown():
        task = getattr(app.state, "_telemetry_task", None)
        if task:
            task.cancel()

    return app


async def _telemetry_poller(app: FastAPI):
    """A simple background task to poll node telemetry periodically and cache it."""
    try:
        while True:
            try:
                node_metrics = app.state.web3.fetch_basic_metrics()
                # save into state cache (simple example)
                app.state.latest_node_metrics = node_metrics
            except Exception:
                # swallow errors to avoid crashing background loop
                app.state.latest_node_metrics = {"error": "unreachable"}
            await asyncio.sleep(settings.TELEMETRY_POLL_INTERVAL)
    except asyncio.CancelledError:
        return
