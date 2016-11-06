import requests, sys, threading
from config import reciving_server, timimg_config

def post_volumes(endpoint, data, headers, timestamp):
    d = threading.Thread(name="Volume post: " + str(timestamp), target=excecute_post, args=(endpoint, data, headers))
    d.setDaemon(True)
    d.start()
    pass

def excecute_post(endpoint, data, headers):
    try:
        requests.post(url=endpoint, data=data, headers=headers, timeout=timimg_config["TIMEOUT"])
    except requests.ConnectionError as e:
        print "Connection error, check status of endpoint"
    except requests.RequestException as e:
        print e
        sys.exit(1)
