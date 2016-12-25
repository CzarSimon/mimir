from tweepy import OAuthHandler
from tweepy import Stream
from tweepy.streaming import StreamListener
import sys, logging, traceback, time
from datetime import datetime
import stocks
from config import twitter_credentials, timing

logging.basicConfig(filename="miner.log", level=logging.ERROR)

class MyListener(StreamListener):
    def __init__(self, tickers, aliases, stock_querys):
        self.tickers = tickers
        self.aliases = aliases
        self.stock_querys = stock_querys

    def on_data(self, data):
        print '----- New tweet -----'
        logging.info("New tweet added at: " + _utcnow())
        stocks.store_tweet(data, self.tickers, self.aliases)
        stocks.send_url_to_ranker(data, self.stock_querys)
        return True

    def on_error(self, status_code):
        print "Tweepy error code: {}".format(status_code)
        _check_for_rate_limit(status_code)

def _check_for_rate_limit(status_code):
    if status_code == 420:
        print "Rate limit response. Pausing for {} seconds".format(timing["RATE_LIMIT_PAUSE"])
        time.sleep(timing["RATE_LIMIT_PAUSE"])

def _get_twtr_auth(credentials):
    auth = OAuthHandler(credentials["consumer_key"], credentials["consumer_secret"])
    auth.set_access_token(credentials["access_token"], credentials["access_secret"])
    return auth

def _get_listener():
    ticker_list, stock_querys = stocks.get_stocks_info()
    tickers = set(ticker_list)
    aliases = stocks.getAliases()
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
        else:
            stocks.getStockTweets()
    return planned_exit

if __name__ == "__main__":
    planned_exit = False
    while not planned_exit:
        planned_exit = main()
    print "Terminated by user command"
