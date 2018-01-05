# standard library
import json
import time
import threading
import requests

# user specific import
import communication as comm
import db
import util


# get_names_and_tickers Retrives tickers and names of stocks to track
def get_names_and_tickers():
    query = 'SELECT TICKER, NAME FROM STOCKS WHERE IS_TRACKED=TRUE'
    results = db.query_db(query)
    tickers = parse_tickers(results)
    stock_info = parse_stock_info(results)
    return tickers, stock_info


# parse_tickers Parses ticker name from database result
def parse_tickers(stocks):
    return map(lambda stock: _add_dollar_tag(str(stock[0])), stocks)


# parse_stock_info Parses stocks ticker and name into a dictionary with ticker
#                  as key and a dictonary of ticker and name as value
def parse_stock_info(stocks):
    stock_info = {}
    for stock in stocks:
        ticker = str(stock[0])
        name = str(stock[1])
        stock_info[ticker] = dict(
            ticker=ticker,
            name=name
        )
    return stock_info


# print_tweet_count Prints the number of tweets that are stored
def print_tweet_count():
    tweet_count = db.query_db('SELECT count(*) FROM stockTweets', False)[0]
    pretty_count = "{:,}".format(tweet_count).replace(",", " ")
    print "Tweets stored: {}".format(pretty_count)


# get_aliases Gets all the ticker aliases from stored in the database
def get_aliases():
    aliases = db.query_db('SELECT alias, ticker FROM tickerAliases', True)
    aliDict = {}
    for alias in aliases:
        aliDict[str(alias[0])] = str(alias[1])
    return aliDict


# _add_dollar_tag Prepends a ticker with a dollar tag
def _add_dollar_tag(ticker):
    return "${}".format(ticker)


# _remove_dollar_tag Strips a leading dollar tag from a ticker
def _remove_dollar_tag(ticker):
    return ticker.replace("$", "")


# record_untracked Parses untracked ticker and recordes its occurance
def record_untracked(tweet_id, ticker, timestamp):
    clean_ticker = _remove_dollar_tag(ticker)
    print "Recording occurance of ticker: {}".format(clean_ticker)
    db.insert_untracked(tweet_id, clean_ticker, timestamp)


# store_tweet Parses and inserts a tweet in the database
def store_tweet(tweet, tickers, aliases):
    tweet_text = tweet['text']
    print tweet_text
    created_at = _parse_tweet_time(tweet['created_at']) # Converting tweet time format to date format
    symbols = tweet['entities']['symbols']
    check_and_store(tweet['id_str'], tweet["user"]["id_str"], created_at, tweet_text, tweet['lang'], tweet['user']['followers_count'], symbols, tickers, aliases)


# store_tweet_and_tickers Saves tweets and untracked tickers in the database
def store_tweet_and_tickers(tweets, untracked_tickers):
    if len(tweets) > 0:
        for tweet in tweets:
            db.insert_tweet(tweet)
        tweet_id = tweets[0]['tweet_id']
        now = util.utcnow()
        for ticker in untracked_tickers:
            record_untracked(tweet_id, ticker, now)


# _parse_tweet_time Returns formated created at time from tweet
def _parse_tweet_time(timestamp):
    time.strftime('%Y-%m-%d %H:%M:%S', time.strptime(timestamp,'%a %b %d %H:%M:%S +0000 %Y'))


# check_and_store Checks if ticker is tracked and either stores as tweet or records untraced occurance
def check_and_store(tweet_id, userId, date, tweet, lang, followers, symbols, tickers, aliases):
    upper_symbols = map(lambda symbol: symbol["text"].upper(), symbols)
    unique_symbols = list(set(upper_symbols)) # pretty sure the casting back to list is unnececary
    for symbol in unique_symbols:
        sym = check_ticker(symbol, tickers, aliases)
        if sym['success']:
            db.insert_tweet(tweet_id, userId, date, tweet, sym["ticker"], lang, followers)
        else:
            record_untracked(tweet_id, sym["ticker"], date)


# check_ticker Checks if ticker is tracked, either as a plain ticker or as an alias
def check_ticker(ticker, tickers, aliases):
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
