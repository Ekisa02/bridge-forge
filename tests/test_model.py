import torch
from src.nebula_bridge.models.torch_model import TransactionClassifier


def test_model_predict_shapes():
    model = TransactionClassifier(input_dim=4)
    label, score = model.predict([0.1, 0.2, -0.1, 0.0])
    assert isinstance(label, str)
    assert isinstance(score, float)
    assert 0.0 <= score <= 1.0


def test_save_and_load(tmp_path):
    model = TransactionClassifier(input_dim=4)
    p = tmp_path / "ckpt.pt"
    model.save(str(p))
    loaded = TransactionClassifier.load_from_checkpoint(str(p))
    # Make sure predict still runs
    label, score = loaded.predict([0.0, 0.0, 0.0, 0.1])
    assert isinstance(label, str)
