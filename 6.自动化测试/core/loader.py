import os 
import yaml 

def load_yaml(rel_path: str):
    root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    path = os.path.join(root,"data",rel_path)
    with open(path,encoding="utf-8") as f:
        return yaml.safe_load(f)