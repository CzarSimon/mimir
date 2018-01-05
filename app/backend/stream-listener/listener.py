# standard library
import time

# external library
from tweepy.streaming import StreamListener

# user specific inport
import tweet
from config import timing
import util


# MimirListener Listner class for twitter stream of ticker tweets
class MimirListener(StreamListener):
    def __init__(self, tickers, aliases, stock_querys):
        self.error_count = 0
        self.tracking_data = dict(
            tickers=tickers,
            aliases=aliases,
            stock_querys=stock_querys
        )

    # on_data Method invoced on new tweet
    def on_data(self, data):
        self.error_count = 0
        print '----- New tweet -----'
        tweet.thread_tweet_actions(data, self.tracking_data)
        return True

    # on_error Method invoced to handle errors
    def on_error(self, status_code):
        self.error_count += 1
        print "Tweepy error code: {}".format(status_code)
        handle_twitter_error(status_code, self.error_count)


# handle_twitter_error Handles errors comming form the stream listener
def handle_twitter_error(status_code, error_count):
    pause_seconds = timing["TWITTER_ERROR_PAUSE"] * error_count
    RATE_LIMIT_CODE = 420
    if status_code == RATE_LIMIT_CODE:
        print "Rate limit response. Pausing for {} seconds".format(timing["TWITTER_ERROR_PAUSE"])
        pause_seconds *= 2
    time.sleep(pause_seconds)
