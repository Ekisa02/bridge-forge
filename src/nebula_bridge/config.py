import os
from dataclasses import dataclass

from dotenv import load_dotenv

load_dotenv()


@dataclass
class Settings:
    WEB3_PROVIDER_URL: str = os.getenv("WEB3_PROVIDER_URL", "")
    TELEMETRY_POLL_INTERVAL: int = int(os.getenv("TELEMETRY_POLL_INTERVAL", "60"))
    MODEL_PATH: str = os.getenv("MODEL_PATH", "models/checkpoint.pt")


settings = Settings()
