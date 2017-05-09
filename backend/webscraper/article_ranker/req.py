import requests
import sys


def send_response(data):
    endpoint = "http://localhost:5000/ranked-article"
    try:
        requests.post(url=endpoint, data=data, headers=make_headers(), timeout=0.5)
    except requests.ConnectionError as e:
        sys.stderr.write("Connection error, check status of endpoint")
    except requests.RequestException as e:
        sys.stderr.write(e)


def make_headers():
    return {'content-type': 'application/json'}
