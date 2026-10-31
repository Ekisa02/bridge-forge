# Nebula Bridge Architecture

Nebula Bridge is intentionally small and extensible. Key modules:

- src/nebula_bridge/app.py: Application factory, startup/shutdown handlers, background poller
- src/nebula_bridge/api/v1: FastAPI routers (health, telemetry, classification)
- src/nebula_bridge/models: PyTorch model implementation and utilities
- src/nebula_bridge/svc: service wrappers, e.g., Web3 provider interactions
- scripts: convenience scripts (train demo model)

Design goals
- Keep adapters (web3 clients) small and replaceable
- Keep ML model interface consistent: save/load/predict
- Use FastAPI startup events to initialize expensive resources

Telemetry flow
- Background poller fetches metrics from provider and caches them in app.state
- /v1/telemetry/node returns cached metrics or falls back to live fetch

Classification flow
- /v1/classify accepts a JSON array of features
- Model's predict() returns a label + confidence
- Model checkpoint can be provided via MODEL_PATH env var
