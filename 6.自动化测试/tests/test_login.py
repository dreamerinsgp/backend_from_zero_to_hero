import pytest
from core.loader import load_yaml 

def _load_users():
    return load_yaml("users.yaml")

@pytest.mark.smoke
@pytest.mark.regression
@pytest.mark.parametrize("case",_load_users())
def test_login_parametrized(client, base_url,case):
    r = client.post(
        f"{base_url}/login",
        json={"email": case["email"], "password": case["password"]},
    )
    assert r.status_code == case["expect_status"]


@pytest.mark.regression 
def test_login_missing_email(client,base_url):
    r = client.post(f"{base_url}/login",json={"password":"x"})
    assert r.status_code == 400 

@pytest.mark.regression
def test_login_invalid_json(client,base_url):
    r = client.post(f"{base_url}/login",data="not json",headers={"Content-Type": "application/json"},)
    assert r.status_code == 400 
