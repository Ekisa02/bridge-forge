from typing import List
from fastapi import APIRouter, Request, HTTPException
from pydantic import BaseModel

router = APIRouter()


class ClassifyRequest(BaseModel):
    features: List[float]


class ClassifyResponse(BaseModel):
    label: str
    score: float


@router.post("/classify", response_model=ClassifyResponse)
async def classify(req: ClassifyRequest, request: Request):
    model = getattr(request.app.state, "model", None)
    if model is None:
        raise HTTPException(status_code=503, detail="model not loaded")

    # model expects a 1D list of floats
    try:
        label, score = model.predict(req.features)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))

    return {"label": label, "score": float(score)}
