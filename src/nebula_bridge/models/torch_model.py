from typing import Tuple, Optional
import torch
import torch.nn as nn
import torch.nn.functional as F
import os


class TransactionClassifier(nn.Module):
    def __init__(self, input_dim: int = 4, hidden: int = 32, num_classes: int = 2):
        super().__init__()
        self.fc1 = nn.Linear(input_dim, hidden)
        self.fc2 = nn.Linear(hidden, hidden)
        self.out = nn.Linear(hidden, num_classes)
        self.labels = ["benign", "suspicious"]
        self.input_dim = input_dim

    def forward(self, x):
        x = F.relu(self.fc1(x))
        x = F.relu(self.fc2(x))
        return self.out(x)

    def predict(self, features) -> Tuple[str, float]:
        # features: List[float]
        if len(features) != self.input_dim:
            raise ValueError(f"expected {self.input_dim} features, got {len(features)}")
        self.eval()
        with torch.no_grad():
            x = torch.tensor(features, dtype=torch.float32).unsqueeze(0)
            logits = self.forward(x)
            probs = F.softmax(logits, dim=1).squeeze(0)
            score, idx = torch.max(probs, dim=0)
            return self.labels[int(idx)], float(score)

    def save(self, path: str):
        os.makedirs(os.path.dirname(path), exist_ok=True)
        torch.save({
            "state_dict": self.state_dict(),
            "input_dim": self.input_dim,
        }, path)

    @classmethod
    def load_from_checkpoint(cls, path: str, input_dim: int = 4, device: Optional[str] = None):
        # If checkpoint is missing, return a fresh model with default random weights
        model = cls(input_dim=input_dim)
        if not path:
            return model
        try:
            ckpt = torch.load(path, map_location=device or "cpu")
            model.load_state_dict(ckpt.get("state_dict", ckpt))
            return model
        except Exception:
            # fallback to new model
            return model
