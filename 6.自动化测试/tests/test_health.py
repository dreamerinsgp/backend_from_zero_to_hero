import pytest
import requests 

from core.config import BASE_URL,TIMEOUT

@pytest.mark.smoke
@pytest.mark.regression
def test_health_returns_ok(client,base_url):
    r = client.get(f"{base_url}/health")
    assert r.status_code == 200
    j = r.json()
    assert j.get("status") == "ok"

