import requests, threading
from config import timing

def post_request(endpoint, data, headers, identifier, is_threaded):
    if is_threaded:
        d = threading.Thread(name=identifier, target=_excecute_post, args=(endpoint, data, headers, identifier))
        d.setDaemon(True)
        d.start()
    else:
        _excecute_post(endpoint, data, headers, identifier)

def _excecute_post(endpoint, data, headers, identifier):
    try:
        requests.post(url=endpoint, data=data, headers=headers, timeout=timing["TIMEOUT"])
    except requests.ConnectionError as e:
        print "Connection error caused in {}. Check endpoint status".format(identifier)
    except requests.RequestException as e:
        print e
