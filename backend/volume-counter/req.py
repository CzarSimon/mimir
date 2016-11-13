import requests, sys, threading
from config import reciving_server, timimg

def post_volumes(endpoint, data, headers, timestamp):
    d = threading.Thread(name="Volume post: " + str(timestamp), target=_excecute_post, args=(endpoint, data, headers))
    d.setDaemon(True)
    d.start()
    pass

def _excecute_post(endpoint, data, headers):
    try:
        requests.post(url=endpoint, data=data, headers=headers, timeout=timimg["TIMEOUT"])
    except requests.ConnectionError as e:
        print "Connection error, check status of endpoint"
    except requests.RequestException as e:
        print e
        sys.exit(1)
