import sys
import json
import time
import threading
sys.path.append("..")

from database import manager as db
from urgencyModule import urgency

def dbSetup(do):
    if do:
        print "Press 1 - to reset database"
        reset = input()
        if reset == 1:
            print "Setting up database"
            db.setupTables(do)
            return True
    print "Not setting up db"

def getStockTickers():
    dbTickers = db.queryDatabase('SELECT ticker FROM stocks ORDER BY name', True)
    tickers = []
    for ticker in dbTickers:
        tickers += [ticker[0]]
    return tickers

def getStockTweets(returnList=False):
    if not returnList:
        stockTweets = db.queryDatabase('SELECT ticker, tweet FROM stockTweets ORDER BY ticker', True)
        print "There are " + str(len(stockTweets)) + " stored tweets in the database"
    else:
        stockTweets = db.queryDatabase('SELECT ticker, createdAt FROM stockTweets ORDER BY ticker', True)
        return stockTweets

def getAliases():
    aliases = db.queryDatabase('SELECT alias, ticker FROM tickerAliases', True)
    aliDict = {}
    for alias in aliases:
        aliDict[str(alias[0])] = str(alias[1])
    return aliDict

def addStock(ticker):
    print 'Adding stock with ticker: ' + ticker

def addAlias(alias, ticker):
    db.addAlias(alias, ticker)
    return True

def storeTweet(data, tickers, aliases, report):
    tweet = json.loads(data)
    tweet_text = tweet['text'].encode('utf-8')
    print tweet_text
    expandedUrls = []
    urls = tweet['entities']['urls']
    for url in urls:
        expandedUrls += [str(url['expanded_url'])]
    urls = str(expandedUrls)
    createdAt = time.strftime('%Y-%m-%d %H:%M:%S', time.strptime(tweet['created_at'],'%a %b %d %H:%M:%S +0000 %Y')) # Converting tweet time format to date format
    symbols = tweet['entities']['symbols']
    updatedVolumeTickers = storeTickerTweets(tweet['id_str'], tweet["user"]["id_str"], createdAt, tweet_text, urls, tweet['lang'], tweet['user']['followers_count'], symbols, tickers, aliases, report)
    return updatedVolumeTickers

def storeTickerTweets(tweetId, userId, date, tweet, urls, lang, followers, symbols, tickers, aliases, report):
    # Add check for unique tickers
    upper_symbols = []
    for symbol in symbols:
        upper_symbols += ["$" + symbol['text'].upper()]
    unique_symbols = list(set(upper_symbols))
    for symbol in unique_symbols:
        sym = checkTicker(symbol, tickers, aliases)
        if sym['success']:
            tickers[sym["ticker"]] += 1
            db.insertTweet(tweetId, userId, date, tweet, sym["ticker"], urls, lang, followers)
        else:
            addStock(sym["ticker"])
    if report:
        threadedUrgency(tickers)
    return tickers

def threadedUrgency(tickers):
    daemonThread = threading.Thread(name="rank urgency thread", target=urgency.rankUrgency, args=(tickers,))
    daemonThread.setDaemon(True)
    daemonThread.start()
    return True

def threadedInsert(tweetId, userId, date, tweet, urls, lang, followers, symbols, tickers, aliases):
    d = threading.Thread(name="Tweet insert thread", target=storeTickerTweets, args=(tweetId, userId, date, tweet, urls, lang, followers, symbols, tickers, aliases,))
    d.setDaemon(True)
    d.start()
    return True

def checkTicker(candidate, tickers, aliases):
    if candidate in tickers:
        print "In filter: " + candidate
        return {"success": True, "ticker": candidate}
    elif candidate in aliases:
        print "In aliases: " + candidate + " treated as " + aliases[candidate]
        return {"success": True, "ticker": aliases[candidate]}
    else:
        return {"success": False, "ticker": candidate}

