from fastapi import APIRouter, Request

router = APIRouter()

@router.get("/node")
async def node_telemetry(request: Request):
    """Return latest node telemetry as polled by background task or fetch on demand."""
    # Prefer cached metrics from background poll
    metrics = getattr(request.app.state, "latest_node_metrics", None)
    if metrics:
        return {"cached": True, "metrics": metrics}

    # Fallback: fetch live
    web3 = getattr(request.app.state, "web3", None)
    if web3:
        try:
            metrics = web3.fetch_basic_metrics()
            return {"cached": False, "metrics": metrics}
        except Exception as e:
            return {"error": str(e)}

    return {"error": "no web3 provider configured"}
