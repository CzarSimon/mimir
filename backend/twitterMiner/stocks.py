import sys, json, time, threading, requests, communication as comm
from urlparse import urlparse
from config import APP_SERVER, NEWS_SERVER, forbidden_domains
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

def get_stocks_info():
    url = "".join([APP_SERVER["ADDRESS"], APP_SERVER["routes"]["STOCKLIST"]])
    stock_querys = {}
    try:
        res = requests.get(url)
        stock_list = json.loads(res.content)["list"]
        ticker_list = map(lambda stock: "$" + str(stock["ticker"]), stock_list)
        for stock in stock_list:
            stock_querys[str(stock["ticker"])] = {"ticker": str(stock["ticker"]), "name": str(stock["name"])}
        return {"success": True, "ticker_list": ticker_list, "stock_querys": stock_querys}
    except Exception as e:
        print e
        return {"success": False, "error": e}


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

def send_url_to_ranker(data, stock_querys):
    tweet = json.loads(data)
    entities = tweet["entities"]
    urls = map(lambda url: url["expanded_url"], entities["urls"])
    filtered_urls = filter(lambda url: (url is not None) and (urlparse(url).netloc not in forbidden_domains), urls)
    if (len(filtered_urls) > 0):
        rank_object = _create_rank_object(stock_querys, filtered_urls, entities["symbols"], tweet["user"], tweet["lang"])
        if (_control_rank_object(rank_object)):
            rank_url = "".join([NEWS_SERVER["ADDRESS"], NEWS_SERVER["routes"]["RANK"]])
            comm.post_request(rank_url, json.dumps(rank_object), {'content-type': 'application/json'}, "send_url_to_ranker", is_threaded=True)
        else:
            pass
    else:
        pass

def _create_rank_object(stock_querys, urls, symbols, author, lang):
    tickers = map(lambda symbol: symbol["text"].upper(), symbols)
    relevant_tickers = filter(lambda ticker: ticker in stock_querys, tickers)
    return {
        "urls": urls,
        "subjects": map(lambda ticker: stock_querys[ticker], relevant_tickers),
        "author": {"id": author["id"], "follower_count": author["followers_count"]},
        "language": lang
    }

def _control_rank_object(rank_obj):
    if (len(rank_obj["urls"]) < 1):
        return False
    elif (len(rank_obj["subjects"]) < 1):
        return False
    elif ("id" not in rank_obj["author"] and "follower_count" not in rank_obj["author"]):
        return False
    else:
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
