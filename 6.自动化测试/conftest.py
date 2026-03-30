#作用：fixture,hooks,全局变量，pytest自动发现及加载

import os
import sys 
import pytest
# pytest 默认会把 tests/ 所在目录加入 sys.path，但有时（例如从上级目录运行、或 CI 环境）可能不包含项目根，导致 import core 失败。
# 在 conftest.py 里显式插入项目根，可以保证无论从哪里执行 pytest，都能正确导入 core 包。

sys.path.insert(0,os.path.dirname(os.path.abspath(__file__)))

from core.client import get_client 
from core.loader import load_yaml 

@pytest.fixture(scope="session")
def base_url():
    return os.getenv("BASE_URL","http://localhost:5000")

# 通过参数 base_url 注入，Pytest 会先执行 base_url 再执行 client
@pytest.fixture
def client(base_url):
    return get_client(base_url)

@pytest.fixture(scope="session")
def auth_token(base_url):
    users = load_yaml("users.yaml")
    cred = next(u for u in users if u.get("expect_status")==200)
    c = get_client(base_url)
    r = c.post(f"{base_url}/login", json={"email":cred["email"],"password": cred["password"]},)
    assert r.status_code == 200 , f"login failed: {r.text}"
    data = r.json()
    assert "token" in data 
    return data["token"]