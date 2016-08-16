from tweepy import OAuthHandler
from tweepy import Stream
from tweepy.streaming import StreamListener
import sys
import logging, schedule, threading, time
from datetime import datetime

sys.path.append("..")
import twitterCredentials
import stocks
from urgencyModule import meanAndStdevCalc as calc

consumerKey = twitterCredentials.consumerKey
consumerSecret = twitterCredentials.consumerSecret
accessToken = twitterCredentials.accessToken
accessSecret = twitterCredentials.accessSecret

logging.basicConfig(filename="miner.log", level=logging.ERROR)

class MyListener(StreamListener):
    def __init__(self, tickers, aliases, stock_querys):
        self.tickers = tickers
        self.aliases = aliases
        self.stock_querys = stock_querys
        self.lastReportMinute = -1

    def on_data(self, data):
        minute = datetime.utcnow().minute
        if minute < self.lastReportMinute:
            for ticker in self.tickers.iteritems():
                self.tickers[ticker] = 0
            report = True
            self.lastReportMinute = minute
        elif self.lastReportMinute < 0 or (self.lastReportMinute+1 <= minute):
            report = True
            self.lastReportMinute = minute
        else:
            report = False
        print '----- New tweet -----'
        logging.info("New tweet added at: " + str(datetime.utcnow()))
        self.tickers = stocks.storeTweet(data, self.tickers, self.aliases, report)
        stocks.send_url_to_ranker(data, self.stock_querys)
        return True

    def on_error(self, status_code):
        print (status_code)
        return True


def list2Dict(list):
    dict = {}
    for item in list:
        dict[str(item)] = 0
    return dict

def doMeanAndStdevCalc():
    schedule.every().day.at('23:50').do(calc.getStockTweets) # Adjust this run to time zone
    while True:
        schedule.run_pending()
        time.sleep(60)

def main():
    print "Runnig"
    auth = OAuthHandler(consumerKey, consumerSecret)
    auth.set_access_token(accessToken, accessSecret)
    stocks_info = stocks.get_stocks_info()
    tickers = stocks_info["ticker_list"]
    tickerDict = list2Dict(tickers)
    tickerAliases = stocks.getAliases()
    plannedExit = False

    try:
        twitterStream = Stream(auth, MyListener(tickerDict, tickerAliases, stocks_info["stock_querys"]))
        twitterStream.filter(track=tickers)
    except (KeyboardInterrupt, Exception) as e:
        if str(type(e)) == "<type 'exceptions.KeyboardInterrupt'>":
            plannedExit = True
        else:
            print "This was the error: " + str(e)
            logging.error("In twitterMiner.main() - " + str(e) + " - " + str(datetime.utcnow()))
            twitterStream.disconnect()
    finally:
        if not plannedExit:
            errorStr = "Twitter miner ended unexpectedly at: " + str(datetime.utcnow())
            print errorStr
            logging.debug(errorStr)
        else:
            stocks.getStockTweets()
        return plannedExit


def test_run():
    print "Starting test"
    result = stocks.get_stocks_info()
    print result["ticker_list"]
    print result["stock_querys"]
    print "Done"

if __name__ == "__main__":
    # test_run()
    d = threading.Thread(name="caluclation thread", target=doMeanAndStdevCalc)
    d.setDaemon(True)
    d.start()
    plannedExit = False
    while not plannedExit:
        plannedExit = main()
    print "Terminated by user command"