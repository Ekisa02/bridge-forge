import pytest
from httpx import AsyncClient
from src.nebula_bridge.main import app


@pytest.mark.asyncio
async def test_health():
    async with AsyncClient(app=app, base_url="http://test") as ac:
        r = await ac.get("/v1/health")
        assert r.status_code == 200
        assert r.json()["status"] == "ok"


@pytest.mark.asyncio
async def test_classify_default_model():
    # call classify with 4 features against the bundled model (random weights if none)
    async with AsyncClient(app=app, base_url="http://test") as ac:
        payload = {"features": [0.1, 0.2, -0.3, 0.5]}
        r = await ac.post("/v1/classify", json=payload)
        assert r.status_code == 200
        body = r.json()
        assert "label" in body and "score" in body
        assert 0.0 <= body["score"] <= 1.0
