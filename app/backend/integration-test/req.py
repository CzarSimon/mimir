import json
import requests


def post_data(data, url):
    headers = {'content-type': 'application/json'}
    res = requests.post(url=url, data=json.dumps(data), headers=headers)
    return res
