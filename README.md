# Nebula Bridge

Nebula Bridge is a modular Web3 infrastructure toolkit built with FastAPI and PyTorch. It provides components for node telemetry, transaction classification with a pluggable ML model, and adapter-style integrations for plugging different Web3 providers.

Key features
- FastAPI-based REST API for telemetry, classification, and health checks
- Lightweight PyTorch transaction classifier with save/load/predict utilities
- Web3 provider adapter for Node metrics and health
- Background telemetry polling and in-memory caching
- Example training script to produce a demo model checkpoint
- Tests using pytest and httpx

Architecture

ASCII diagram:

  +----------------------+      +----------------------+      +----------------+
  |   HTTP Clients       | ---> |   FastAPI App        | ---> | PyTorch Model  |
  | (CLI / Grafana / UI) |      | (routers, services)  |      | (transaction    |
  +----------------------+      +----------------------+      |  classifier)    |
                                       |  ^   ^                +----------------+
                                       |  |   |
                             +---------+  |   +----------+
                             |            |              |
                    +--------v-----+ +----v-------+ +----v-------+
                    | web3_client  | | cache/core | | background  |
                    | (provider)   | | utils      | | tasks       |
                    +--------------+ +------------+ +-------------+


Installation

Requirements: Python 3.9+ (3.10/3.11 recommended)

1. Clone

   git clone https://github.com/your-org/nebula-bridge.git
   cd nebula-bridge

2. Setup virtualenv

   python -m venv .venv
   source .venv/bin/activate

3. Install

   pip install -r requirements.txt

4. Run the app (development)

   # from repo root
   uvicorn nebula_bridge.main:app --reload --port 8000

5. Endpoints

   GET  /health                     -> service health
   GET  /v1/telemetry/node          -> node telemetry (queries web3 provider when configured)
   POST /v1/classify                -> classify transaction features (JSON payload)

Example classify request

  POST /v1/classify
  {
    "features": [0.1, 0.3, 0.2, 0.4]
  }

Response:
  {
    "label": "benign",
    "score": 0.87
  }

Model

A small MLP model implementation lives in src/nebula_bridge/models/torch_model.py. You can create a demo checkpoint using:

  python scripts/train_demo_model.py --out models/checkpoint.pt

If no checkpoint is available the API will instantiate a fresh model with random weights (useful for development).

Configuration

Configuration follows environment variables (see src/nebula_bridge/config.py). Create a .env with keys such as:

  WEB3_PROVIDER_URL=https://mainnet.infura.io/v3/YOURKEY
  TELEMETRY_POLL_INTERVAL=60

Testing

  pip install -r requirements.txt
  pytest -q

Docker

A sample Dockerfile is provided. Build and run:

  docker build -t nebula-bridge:dev .
  docker run -p 8000:8000 nebula-bridge:dev

Contributing

Contributions welcome. Structure follows a modular design inside src/nebula_bridge for easy extension. Please run tests and linters before opening PRs.

License

MIT
