Model checkpoint directory.

Run scripts/train_demo_model.py to create a demo checkpoint for the small MLP used by the classifier:

  python scripts/train_demo_model.py --out models/checkpoint.pt --epochs 20

If a checkpoint does not exist the API will instantiate a fresh model with random weights.
