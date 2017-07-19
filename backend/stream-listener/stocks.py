# standard library
import json
import time
import threading
import requests

# user specific import
import communication as comm
from config import APP_SERVER
import db


def get_stocks_info():
    url = comm.to_url(APP_SERVER, "STOCKLIST")
    stock_querys = {}
    res = requests.get(url)
    stock_list = json.loads(res.content)["list"]
    ticker_list = map(lambda stock: "$" + str(stock["ticker"]), stock_list)
    for stock in stock_list:
        stock_querys[str(stock["ticker"])] = {"ticker": str(stock["ticker"]), "name": str(stock["name"])}
    return ticker_list, stock_querys


def print_tweet_count():
    tweet_count = db.queryDatabase('SELECT count(*) FROM stockTweets', False)[0]
    pretty_count = "{:,}".format(tweet_count).replace(",", " ")
    print "Tweets stored: {}".format(pretty_count)


def getAliases():
    aliases = db.queryDatabase('SELECT alias, ticker FROM tickerAliases', True)
    aliDict = {}
    for alias in aliases:
        aliDict[_add_dollar_tag(str(alias[0]))] = _add_dollar_tag(str(alias[1]))
    return aliDict


def _add_dollar_tag(ticker):
    return "${}".format(ticker)


def _remove_dollar_tag(ticker):
    return ticker.replace("$", "")


def record_untracked(tweet_id, ticker, timestamp):
    clean_ticker = ticker.replace("$", "")
    print "Recording occurance of ticker: {}".format(clean_ticker)
    db.insert_untracked(tweet_id, clean_ticker, timestamp)


def store_tweet(tweet, tickers, aliases):
    tweet_text = tweet['text']
    print tweet_text
    created_at = _parse_tweet_time(tweet['created_at']) # Converting tweet time format to date format
    symbols = tweet['entities']['symbols']
    check_and_store(tweet['id_str'], tweet["user"]["id_str"], created_at, tweet_text, tweet['lang'], tweet['user']['followers_count'], symbols, tickers, aliases)


# _parse_tweet_time Returns formated created at time from tweet
def _parse_tweet_time(timestamp):
    time.strftime('%Y-%m-%d %H:%M:%S', time.strptime(timestamp,'%a %b %d %H:%M:%S +0000 %Y'))


def check_and_store(tweet_id, userId, date, tweet, lang, followers, symbols, tickers, aliases):
    upper_symbols = map(lambda symbol: symbol["text"].upper(), symbols)
    unique_symbols = list(set(upper_symbols)) # pretty sure the casting back to list is unnececary
    for symbol in unique_symbols:
        sym = checkTicker(symbol, tickers, aliases)
        if sym['success']:
            db.insert_tweet(tweet_id, userId, date, tweet, sym["ticker"], lang, followers)
        else:
            record_untracked(tweet_id, sym["ticker"], date)


def threadedInsert(tweetId, userId, date, tweet, lang, followers, symbols, tickers, aliases):
    d = threading.Thread(name="Tweet insert thread", target=storeTickerTweets, args=(tweetId, userId, date, tweet, lang, followers, symbols, tickers, aliases,))
    d.setDaemon(True)
    d.start()
    return True


def checkTicker(ticker, tickers, aliases):
    candidate = _add_dollar_tag(ticker)
    if candidate in tickers:
        print "In filter: " + ticker
        return {"success": True, "ticker": ticker}
    elif candidate in aliases:
        tickerAlias = _remove_dollar_tag(aliases[candidate])
        print "In aliases: " + ticker + " treated as " + tickerAlias
        return {"success": True, "ticker": tickerAlias}
    else:
        return {"success": False, "ticker": ticker}
