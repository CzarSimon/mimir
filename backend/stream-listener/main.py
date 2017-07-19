# standard library
import sys
import logging
import traceback
import time
from datetime import datetime

# external library
from tweepy import OAuthHandler
from tweepy import Stream
from tweepy.streaming import StreamListener

# user specific inport
import stocks
import tweet
from config import twitter_credentials, timing


logging.basicConfig(filename="miner.log", level=logging.ERROR)


class MyListener(StreamListener):
    def __init__(self, tickers, aliases, stock_querys):
        self.error_count = 0
        self.tracking_data = dict(
            tickers=tickers,
            aliases=aliases,
            stock_querys=stock_querys
        )

    def on_data(self, data):
        self.error_count = 0
        print '----- New tweet -----'
        logging.info("New tweet added at: " + _utcnow())
        tweet.thread_tweet_actions(data, self.tracking_data)
        return True

    def on_error(self, status_code):
        self.error_count += 1
        print "Tweepy error code: {}".format(status_code)
        handle_twitter_error(status_code, self.error_count)


def handle_twitter_error(status_code, error_count):
    pause_seconds = timing["TWITTER_ERROR_PAUSE"] * error_count
    RATE_LIMIT_CODE = 420
    if status_code == RATE_LIMIT_CODE:
        print "Rate limit response. Pausing for {} seconds".format(timing["RATE_LIMIT_PAUSE"])
        pause_seconds *= 2
    time.sleep(pause_seconds)


def _get_twtr_auth(credentials):
    auth = OAuthHandler(credentials["consumer_key"], credentials["consumer_secret"])
    auth.set_access_token(credentials["access_token"], credentials["access_secret"])
    return auth


def _get_listener():
    ticker_list, stock_querys = stocks.get_stocks_info()
    tickers = set(ticker_list)
    aliases = stocks.getAliases()
    print(tickers)
    print(aliases)
    return MyListener(tickers, aliases, stock_querys), ticker_list


def _utcnow():
    return str(datetime.utcnow())


def main():
    print "Runnig"
    planned_exit = False
    listener, ticker_list = _get_listener()
    try:
        twitter_stream = Stream(_get_twtr_auth(twitter_credentials), listener)
        twitter_stream.filter(track=ticker_list)
    except KeyboardInterrupt:
        planned_exit = True
    except Exception as e:
        print "This was the error: " + str(e), type(e)
        traceback.print_exc()
        logging.error(str(e) + " - " + _utcnow())
    finally:
        twitter_stream.disconnect()
        if not planned_exit:
            errorStr = "Twitter miner ended unexpectedly at: " + _utcnow()
            print errorStr
            logging.debug(errorStr)
    return planned_exit


if __name__ == "__main__":
    planned_exit = False
    while not planned_exit:
        planned_exit = main()
        if not planned_exit:
            time.sleep(timing["TWITTER_ERROR_PAUSE"])
    print "Terminated by user command"
    stocks.print_tweet_count()
