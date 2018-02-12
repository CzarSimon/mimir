# standard library
import os


def get_hostname(key):
    return os.getenv(key, "localhost")


def get_service_address(name):
    protocol = os.getenv("{}_PROTOCOL".format(name), "http")
    host = os.getenv("{}_HOST".format(name), "")
    port = os.getenv("{}_PORT".format(name), "")
    return "{}://{}:{}".format(protocol, host, port)


service_name = "Mimir stream listener"


NEWS_SERVER = {
    "ADDRESS": get_service_address("NEWS_RANKER"),
    "routes": {
        "RANK": "/api/rank-article"
    }
}


SPAM_FILTER = {
    "ADDRESS": get_service_address("SPAM_FILTER"),
    "routes": {
        "CLASSIFY": "/classify"
    },
    "handle_spam": os.getenv("HANDLE_SPAM", "FALSE") == "TRUE"
}


timing = {
    "TIMEOUT": 0.5,
    "TWITTER_ERROR_PAUSE": 180
}


database = {
    "NAME": os.getenv("PG_NAME", "mimirprod"),
    "USER": os.getenv("PG_USER","simon"),
    "HOST": get_hostname("PG_HOST"),
    "PASSWORD": os.getenv("PG_PASSWORD", "pwd"),
    "PORT": os.getenv("PG_PORT", "5432")
}


twitter_credentials = {
    "consumer_key": os.environ("TWITTER_CONSUMER_KEY"),
    "consumer_secret": os.environ("TWITTER_CONSUMER_SECRET"),
    "access_token": os.environ("TWITTER_ACCESS_TOKEN"),
    "access_secret": os.environ("TWITTER_ACCESS_TOKEN_SECRET")
}


forbidden_domains = set([
    "owler.us",
    "owler.com",
    "stocktwits.com",
    "investorshangout.com",
    "1broker.com",
    "twitter.com",
    "cityfalcon.com",
    "mixlr.com"
])
