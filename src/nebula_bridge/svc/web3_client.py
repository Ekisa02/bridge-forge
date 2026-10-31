from typing import Dict
from web3 import Web3


class Web3Client:
    def __init__(self, provider_url: str = ""):
        self.provider_url = provider_url
        self.web3 = None
        if provider_url:
            try:
                self.web3 = Web3(Web3.HTTPProvider(provider_url))
            except Exception:
                self.web3 = None

    def is_connected(self) -> bool:
        try:
            return self.web3.is_connected() if self.web3 else False
        except Exception:
            return False

    def fetch_basic_metrics(self) -> Dict:
        # returns a small set of metrics; graceful fallback if provider is not set
        if not self.web3:
            raise RuntimeError("no provider configured")
        try:
            block = self.web3.eth.block_number
            peer_count = getattr(self.web3.net, "peer_count", None)
            syncing = None
            try:
                syncing = self.web3.eth.syncing
            except Exception:
                syncing = False
            return {
                "block_number": block,
                "peer_count": peer_count,
                "syncing": syncing,
                "connected": self.is_connected(),
            }
        except Exception as e:
            raise RuntimeError(f"failed to fetch metrics: {e}")
