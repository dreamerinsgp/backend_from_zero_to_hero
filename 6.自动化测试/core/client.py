import requests 
from .config import BASE_URL, TIMEOUT 

def new_session():
    s = requests.Session()
    s.headers.update({"Content-Type":"application/json"})
    s.timeout = TIMEOUT
    return s 

def get_client(base_url: str = None):
    s = new_session()
    s.base_url = (base_url or BASE_URL).rstrip("/")
    return s