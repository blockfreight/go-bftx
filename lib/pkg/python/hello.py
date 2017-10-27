import json
import requests
from requests.auth import HTTPBasicAuth

def HelloPython():
    url = "http://public.coindaddy.io:4000/api/"
    data = {}
    data['content-type'] = 'application/json'   
    auth = HTTPBasicAuth('rpc', '1234')

    payload = {
    "method": "get_running_info",
    "params": {},
    "jsonrpc": "2.0",
    "id": 0,
    }
    response = requests.post(url, data=json.dumps(payload), headers=data, auth=auth)
    
    return response.text