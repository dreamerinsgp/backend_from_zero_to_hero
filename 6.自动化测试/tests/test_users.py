import pytest

@pytest.mark.regression 
def test_users_list(client,base_url):
    r = client.get(f"{base_url}/users")
    assert r.status_code == 200
    j = r.json()
    assert "users" in j 
    assert len(j["users"]) >= 1

@pytest.mark.regression
@pytest.mark.parametrize("user_id,expect_status",[(1,200),(2,404),(99,404)])
def test_users_get_by_id_parametrized(client,base_url,user_id,expect_status):
    r = client.get(f"{base_url}/users/{user_id}")
    assert r.status_code == expect_status
    if expect_status == 200:
        assert r.json().get("id") == user_id