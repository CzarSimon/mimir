import requests, sys, threading
from config import timimg

def send_stats_to_server(endpoint, data, ticker, headers={'content-type': 'application/json'}):
    d = threading.Thread(name="Sending {} mean & stdev".format(ticker), target=_excecute_post, args=(endpoint, data, headers))
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
