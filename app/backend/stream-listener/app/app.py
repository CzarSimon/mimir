# Standard library
import logging

# 3rd party modules
from tweepy import OAuthHandler, Stream

# Internal modules
from app.config import TwitterConfig, SpamFilterConfig, NewsRankerConfig
from app.mocks import MockFilter, MockRanker
from app.repository import SQLStockRepo, SQLTweetRepo
from app.service import (
    StreamListenerImpl, TweetServiceImpl, SpamFilterService,
    RankingServiceImpl, StreamLogger, FileStreamer)

class App(object):

    __log = logging.getLogger('App')

    def __init__(self):
        self.TRACKED_STOCKS = self.__get_tracked_stocks()
        tweet_svc = self.__setup_tweet_service()
        config = TwitterConfig()
        listener = StreamListenerImpl(tweet_svc, config)
        self.__stream = Stream(self.__setup_auth(config),
                               FileStreamer('/Users/simonlindgren/.mimir/tweets'))

    def start(self):
        cashtags = [f'${symbol}' for symbol in self.TRACKED_STOCKS.keys()]
        self.__log.info(f'Starging with symbols: {cashtags}')
        self.__stream.filter(track=cashtags)

    def __setup_auth(self, config):
        """Sets up auth credentials to listen to twitter stream.

        :param config: TwitterConfig to use for authentication.
        :return: OAuthHandler
        """
        config = TwitterConfig()
        auth = OAuthHandler(config.CONSUMER_KEY, config.CONSUMER_SECRET)
        auth.set_access_token(config.ACCESS_TOKEN, config.ACCESS_TOKEN_SECRET)
        return auth

    def __get_tracked_stocks(self):
        """Get all tracked stocks as a dict with symbols as keys.

        :return: Tracked stocks as dict.
        """
        stock_repo = SQLStockRepo()
        stocks = stock_repo.get_all()
        return {
            s.symbol: {'name': s.name, 'ticker': s.symbol} for s in stocks
        }

    def __setup_tweet_service(self):
        """Sets up a usable tweet service.

        :return: TweetService
        """
        filter_svc = SpamFilterService(SpamFilterConfig())
        ranking_svc = RankingServiceImpl(
            self.TRACKED_STOCKS, NewsRankerConfig())
        tweet_repo = SQLTweetRepo()
        tracked_symbols = set([symbol for symbol in self.TRACKED_STOCKS])
        return TweetServiceImpl(
            tracked_symbols, MockFilter(), MockRanker(), tweet_repo)
