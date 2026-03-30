import pytest 

@pytest.mark.smoke
@pytest.mark.regression
def test_me_with_token(client, base_url, auth_token):
    headers = {"Authorization": f"Bearer {auth_token}"}
    r = client.get(f"{base_url}/me", headers = headers)
    assert r.status_code == 200
    body = r.json()
    assert body.get("id") == 1
    assert "name" in body


def test_me_without_token(client,base_url):
    r = client.get(f"{base_url}/me")
    assert r.status_code == 401