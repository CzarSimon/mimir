# standard library
import json
from threading import Thread

# user specific imports
from config import SPAM_FILTER
from stocks import store_tweet
import communication as comm


# parse_tweet Turns raw tweet into json and returns it
def parse_tweet(raw_tweet):
    return json.loads(raw_tweet)


# thread_tweet_actions Spins of a new thread containg the tweet and tracking_data
def thread_tweet_actions(raw_tweet, tracking_data):
    thread = Thread(target=handle_tweet, args=(raw_tweet, tracking_data))
    thread.setDaemon(True)
    thread.start()
    return


# handle_tweet Parses tweet and sends for storage and ranking
def handle_tweet(raw_tweet, tracking_data):
    tweet = parse_tweet(raw_tweet)
    tweet['text'] = tweet['text'].encode('utf-8')
    if not is_spam(tweet['text']):
        store_tweet(tweet, tracking_data["tickers"], tracking_data["aliases"])
        comm.send_to_ranker(tweet, tracking_data["stock_querys"])
    return


# is_spam Sends tweet test to classifier, returns true is it is classified as spam else not
def is_spam(tweet_text):
    #url = comm.to_url(SPAM_FILTER, "CLASSIFY")
    #response = comm.post_with_response(url, to_classifer_schema(tweet_text), "spam_classify")
    #return response["result"] == "SPAM"
    return False


# to_classifer_schema Constructs a request body compliant with what is expected from the spam-filter 
def to_classifer_schema(tweet_text):
    return {
        "text": tweet_text
    }
