# standard library
import json
import requests
import threading
from urlparse import urlparse

# user specific inport
from config import timing, NEWS_SERVER, forbidden_domains


def threaded_post(endpoint, data, headers, identifier, is_threaded=False):
    d = threading.Thread(name=identifier, target=post, args=(endpoint, data, headers, identifier))
    d.setDaemon(True)
    d.start()


# post Makes a post request to the specified endpoint
def post(endpoint, data, identifier):
    headers = {'content-type': 'application/json'}
    try:
        requests.post(url=endpoint, data=data, headers=headers, timeout=timing["TIMEOUT"])
    except requests.ConnectionError as e:
        print "Connection error caused in {}. Check endpoint status".format(identifier)
    except requests.RequestException as e:
        print e


# post_with_response Makes a post request to the specified endpoint and returns the result
def post_with_response(endpoint, data, identifier):
    response = {"success": True}
    try:
        res = requests.post(url=endpoint, data=data, headers=_get_json_headers(), timeout=timing["TIMEOUT"])
        response["data"] = json.loads(res.content)
    except requests.ConnectionError as e:
        response["success"] = False
        print "Connection error caused in {}. Check endpoint status".format(identifier)
    except requests.RequestException as e:
        response["success"] = False
        print e
    finally:
        return response


def _get_json_headers():
    return {'content-type': 'application/json'}


# to_url Turns a service config into a callable URL
def to_url(service_config, route_key):
    route = service_config["routes"][route_key]
    return service_config["ADDRESS"] + route


# send_to_ranker Creates a rank object from a tweet and sends it to the news ranker
def send_to_ranker(tweet, stock_querys):
    entities = tweet["entities"]
    urls = map(lambda url: _get_url(url), entities["urls"])
    filtered_urls = filter(lambda url: (url is not None) and (urlparse(url).netloc not in forbidden_domains), urls)
    if (len(filtered_urls) > 0):
        rank_object = _create_rank_object(stock_querys, filtered_urls, entities["symbols"], tweet["user"], tweet["lang"])
        if (_control_rank_object(rank_object)):
            rank_url = to_url(NEWS_SERVER, "RANK")
            post(rank_url, json.dumps(rank_object), "send_url_to_ranker")


# _get_url Attempts to get the expanded url from a tweet if it exists
def _get_url(url):
    long_url = url["expanded_url"]
    return long_url if (long_url is not None) else url["url"]


# _create_rank_object Creates the request body to be sent to news ranker
def _create_rank_object(stock_querys, urls, symbols, author, lang):
    tickers = map(lambda symbol: symbol["text"].upper(), symbols)
    relevant_tickers = filter(lambda ticker: ticker in stock_querys, tickers)
    return {
        "urls": urls,
        "subjects": map(lambda ticker: stock_querys[ticker], relevant_tickers),
        "author": {"id": author["id"], "follower_count": author["followers_count"]},
        "language": lang
    }


# _control_rank_object Checks if the rank object contains an author id and
# a follower count plus at least one url and 1 subject
def _control_rank_object(rank_obj):
    if (len(rank_obj["urls"]) < 1):
        return False
    elif (len(rank_obj["subjects"]) < 1):
        return False
    elif ("id" not in rank_obj["author"] and "follower_count" not in rank_obj["author"]):
        return False
    else:
        return True
