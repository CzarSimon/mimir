# standard library
import sys
import traceback
import time
import logging

# external library
from tweepy import OAuthHandler
from tweepy import Stream

# user specific inport
import stocks
import tweet
from config import twitter_credentials, timing
import util
from listener import MimirListener


logging.basicConfig(filename="miner.log", level=logging.ERROR)


# _get_twtr_auth Returns the twitter credentials initialized in a Oauth handler
def _get_twtr_auth(credentials):
    auth = OAuthHandler(credentials["consumer_key"], credentials["consumer_secret"])
    auth.set_access_token(credentials["access_token"], credentials["access_secret"])
    return auth


# _get_listener Retrives stocks to listen for and initializes a stream listener
def _get_listener():
    ticker_list, stock_querys = stocks.get_names_and_tickers()
    tickers = set(ticker_list)
    aliases = stocks.get_aliases()
    print(tickers)
    print(aliases)
    return MimirListener(tickers, aliases, stock_querys), ticker_list


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
        logging.error(str(e) + " - " + util.utcnow())
    finally:
        twitter_stream.disconnect()
        if not planned_exit:
            errorStr = "Twitter miner ended unexpectedly at: " + util.utcnow()
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
