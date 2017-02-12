import json, time, threading, requests, communication as comm
from urlparse import urlparse
from config import APP_SERVER, NEWS_SERVER, forbidden_domains
import db


def get_stocks_info():
    url = "".join([APP_SERVER["ADDRESS"], APP_SERVER["routes"]["STOCKLIST"]])
    stock_querys = {}
    try:
        res = requests.get(url)
        stock_list = json.loads(res.content)["list"]
        ticker_list = map(lambda stock: "$" + str(stock["ticker"]), stock_list)
        for stock in stock_list:
            stock_querys[str(stock["ticker"])] = {"ticker": str(stock["ticker"]), "name": str(stock["name"])}
        return ticker_list, stock_querys
    except Exception as e:
        print e
        return [], {}


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


def _get_url(url):
    long_url = url["expanded_url"]
    return long_url if (long_url is not None) else url["url"]


def send_url_to_ranker(data, stock_querys):
    tweet = json.loads(data)
    entities = tweet["entities"]
    urls = map(lambda url: _get_url(url), entities["urls"])
    filtered_urls = filter(lambda url: (url is not None) and (urlparse(url).netloc not in forbidden_domains), urls)
    if (len(filtered_urls) > 0):
        rank_object = _create_rank_object(stock_querys, filtered_urls, entities["symbols"], tweet["user"], tweet["lang"])
        if (_control_rank_object(rank_object)):
            rank_url = "".join([NEWS_SERVER["ADDRESS"], NEWS_SERVER["routes"]["RANK"]])
            comm.post_request(rank_url, json.dumps(rank_object), {'content-type': 'application/json'}, "send_url_to_ranker", is_threaded=True)


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


def store_tweet(data, tickers, aliases):
    tweet = json.loads(data)
    tweet_text = tweet['text'].encode('utf-8')
    print tweet_text
    created_at = time.strftime('%Y-%m-%d %H:%M:%S', time.strptime(tweet['created_at'],'%a %b %d %H:%M:%S +0000 %Y')) # Converting tweet time format to date format
    symbols = tweet['entities']['symbols']
    check_and_store(tweet['id_str'], tweet["user"]["id_str"], created_at, tweet_text, tweet['lang'], tweet['user']['followers_count'], symbols, tickers, aliases)


def check_and_store(tweet_id, userId, date, tweet, lang, followers, symbols, tickers, aliases):
    upper_symbols = map(lambda symbol: symbol["text"].upper(), symbols)
    unique_symbols = list(set(upper_symbols))
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
