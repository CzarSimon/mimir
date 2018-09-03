# Standard library
import os
from logging.config import dictConfig

# Setup of logging configureaion
from .logging import LOGGING_CONIFG
dictConfig(LOGGING_CONIFG)

from app.config import values
from app.config import util


class DBConfig(object):
    URI = util.get_database_uri()
    ECHO = True


class TwitterConfig(object):
    CONSUMER_KEY = util.getenv("TWITTER_CONSUMER_KEY"),
    CONSUMER_SECRET = util.getenv("TWITTER_CONSUMER_SECRET"),
    ACCESS_TOKEN = util.getenv("TWITTER_ACCESS_TOKEN"),
    TWITTER_ACCESS_TOKEN_SECRET = util.getenv("TWITTER_ACCESS_TOKEN_SECRET")
    ERROR_PAUSE_SECONDS = 180
    TIMEOUT_SECONDS = 0.5


class SpamFilterConfig(object):
    URL = util.getenv('SPAM_FILTER_URL')
    CLASSIFY_ROUTE = '/v1/classify'


class NewsRankerConfig(object):
    URL = util.getenv('NEWS_RANKER_URL')
    RANK_ROUTE = '/v1/article'
