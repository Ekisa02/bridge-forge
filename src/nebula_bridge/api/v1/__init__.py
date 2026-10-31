from fastapi import APIRouter

from .telemetry import router as telemetry_router
from .classify import router as classify_router
from .health import router as health_router

router = APIRouter(prefix="/v1")
router.include_router(telemetry_router, prefix="/telemetry", tags=["telemetry"])
router.include_router(classify_router, prefix="/", tags=["classification"])
router.include_router(health_router, prefix="/", tags=["health"])
