import os 

BASE_URL = os.getenv("BASE_URL","http://localhost:5000")
TIMEOUT = int(os.getenv("TIMEOUT","10"))