# standard library
import json
from threading import Thread
import time
import copy

# user specific imports
from config import SPAM_FILTER
from stocks import store_tweet, store_tweet_and_tickers
import communication as comm
import util


# thread_tweet_actions Spins of a new thread containg the tweet and tracking_data
def thread_tweet_actions(raw_tweet, tracking_data):
    thread = Thread(target=handle_tweet, args=(raw_tweet, tracking_data))
    thread.setDaemon(True)
    thread.start()
    return


# handle_tweet Parses tweet and sends for storage and ranking
def handle_tweet(raw_tweet, tracking_data):
    tweet = parse_tweet(raw_tweet)
    if not is_spam(tweet['text']):
        tracked_tickers, untracked_tickers = filter_tickers(tweet["tickers"], tracking_data)
        tweets = map_tickers_to_tweet(tweet, tracked_tickers)
        _debug_print(tracked_tickers, untracked_tickers, tweets)
        # store_tweet_and_tickers(tweets, untracked_tickers)
        # comm.send_to_ranker(tracked_tickers, tweet, tracking_data["stock_querys"])
    return


def _debug_print(tt, ut, tweets):
    print tt
    print ut
    for tweet in tweets:
        print tweet


# is_spam Sends tweet test to classifier, returns true is it is classified as spam else not
def is_spam(tweet_text):
    if SPAM_FILTER["handle_spam"]:
        return check_for_spam(tweet_text)
    else:
        return False


# check_for_spam Querys the spam classifier if the tweet is spam or not
# retruns true if so, and false if not or connection not successful
def check_for_spam(tweet_text):
    url = comm.to_url(SPAM_FILTER, "CLASSIFY")
    response = comm.post_with_response(url, to_classifer_schema(tweet_text), "spam_classify")
    if response["success"]:
        result = response["data"]["result"]
        print result
        return result == "SPAM"
    else:
        return False


# to_classifer_schema Constructs a request body compliant with what is expected from the spam-filter
def to_classifer_schema(tweet_text):
    return json.dumps(dict(
        text=tweet_text
    ))


# parse_tweet Parses a tweet in json form to a ditctorny containg nececary fields
def parse_tweet(raw_tweet):
    full_tweet = json.loads(raw_tweet)
    return prune_tweet(full_tweet)


# prune_tweet Strips a tweet of unececary fields and formats as nececary
def prune_tweet(tweet):
    #util.pretty_print(tweet)
    entities = tweet["entities"]
    pruned_tweet = dict(
        tweet_id = tweet["id_str"],
        text = parse_tweet_text(tweet["text"]),
        date = parse_tweet_date(tweet["created_at"]),
        tickers = parse_tickers(entities["symbols"]),
        urls = parse_urls(entities["urls"]),
        user_id = tweet["user"]["id_str"],
        user_followers = tweet["user"]["followers_count"],
        lang = tweet["lang"]
    )
    return pruned_tweet


# parse_tweet_date Formats the tweet timestamp if exist, returns current time otherwise
def parse_tweet_date(date_str):
    if date_str is not None and date_str != "":
        ORIGINAL_DATE_FORMAT = '%a %b %d %H:%M:%S +0000 %Y'
        timestamp = time.strptime(date_str, ORIGINAL_DATE_FORMAT)
        TARGET_DATE_FORMAT = '%Y-%m-%d %H:%M:%S'
        return time.strftime(TARGET_DATE_FORMAT, timestamp)
    else:
        return util.utcnow()


# parse_tweet_text Encodes tweet text to be stored in the database
def parse_tweet_text(text):
    return text.encode('utf-8')


# parse_tickers Extracts ticker mentioned in a tweet
def parse_tickers(symbols):
    return map(lambda symbol: symbol["text"].upper(), symbols)


# parse_urls
def parse_urls(urls):
    mapped_urls = map(lambda url: parse_url(url), urls)
    return filter(lambda url: url != "", mapped_urls)


# parse_url Picks url representaion linked in a tweet
def parse_url(url):
    try:
        if "expanded_url" in url and not util.is_empty(url["expanded_url"]):
            return url["expanded_url"]
        else:
            return url["url"]
    except KeyError as e:
        return ""


# map_tickers_to_tweet Creates a list of tweets one for each tracked ticker found
def map_tickers_to_tweet(tweet, tickers):
    tweets = []
    tweet.pop("tickers")
    for ticker in tickers:
        ticker_tweet = copy.deepcopy(tweet)
        ticker_tweet['ticker'] = ticker
        tweets += [ticker_tweet]
    return tweets


# map_aliases Maps a ticker alias to its ticker
def map_aliases(tickers, aliases):
    return map(lambda ticker: aliases[ticker] if ticker in aliases else ticker, tickers)


# filter_tracked Filters and returns only tracked tickers
def filter_tracked(tickers, tracked_tickers):
    return set(filter(lambda ticker: ticker in tracked_tickers, tickers))


# filter_untracked Filters and returns only unctracked tickers
def filter_untracked(tickers, found_tracked_tickers):
    return filter(lambda ticker: ticker not in found_tracked_tickers, tickers)


# filter_tickers Separates tracked and untracted tickers from the observed set
def filter_tickers(tickers, tracking_data):
    mapped_tickers = map_aliases(tickers, tracking_data['aliases'])
    found_tracked_tickers = filter_tracked(mapped_tickers, tracking_data['stock_querys'])
    found_untracked_tickers = filter_untracked(mapped_tickers, found_tracked_tickers)
    return found_tracked_tickers, found_untracked_tickers
