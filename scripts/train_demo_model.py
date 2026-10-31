import argparse
import random
import torch
from src.nebula_bridge.models.torch_model import TransactionClassifier


def synthetic_dataset(n=1000, input_dim=4):
    X = torch.randn(n, input_dim)
    # label heuristic: sum of features > 0 => suspicious else benign
    y = (X.sum(dim=1) > 0).long()
    return X, y


def train(out_path: str, epochs: int = 10, lr: float = 1e-2):
    input_dim = 4
    model = TransactionClassifier(input_dim=input_dim)

    X, y = synthetic_dataset(800, input_dim)
    X_val, y_val = synthetic_dataset(200, input_dim)

    opt = torch.optim.Adam(model.parameters(), lr=lr)
    loss_fn = torch.nn.CrossEntropyLoss()

    model.train()
    for epoch in range(epochs):
        opt.zero_grad()
        logits = model(X)
        loss = loss_fn(logits, y)
        loss.backward()
        opt.step()
        if epoch % 2 == 0:
            with torch.no_grad():
                preds = model(X_val).argmax(dim=1)
                acc = (preds == y_val).float().mean()
            print(f"epoch={epoch}, loss={loss.item():.4f}, val_acc={acc:.4f}")

    model.save(out_path)
    print(f"Saved checkpoint to: {out_path}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--out", dest="out", default="models/checkpoint.pt")
    parser.add_argument("--epochs", dest="epochs", type=int, default=10)
    args = parser.parse_args()
    train(args.out, epochs=args.epochs)
