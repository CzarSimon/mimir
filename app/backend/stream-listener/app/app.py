# Standard library
import logging
from typing import Dict

# 3rd party modules
from tweepy import OAuthHandler, Stream

# Internal modules
from app.config import TwitterConfig, SpamFilterConfig, MQConfig
from app.models import TrackedStock
from app.repository import SQLStockRepo, SQLTweetRepo
from app.service import StreamListenerImpl
from app.service import TweetService, TweetServiceImpl
from app.service import SpamFilterService
from app.service import MQRankingService


class App:

    __log = logging.getLogger('App')

    def __init__(self) -> None:
        self.TRACKED_STOCKS = self.__get_tracked_stocks()
        tweet_svc = self.__setup_tweet_service()
        config = TwitterConfig()
        listener = StreamListenerImpl(tweet_svc, config)
        self.__stream = Stream(self.__setup_auth(config), listener)

    def start(self) -> None:
        cashtags = [f'${symbol}' for symbol in self.TRACKED_STOCKS.keys()]
        self.__log.info(f'Starging with symbols: {cashtags}')
        self.__stream.filter(track=cashtags)

    def __setup_auth(self, config: TwitterConfig) -> OAuthHandler:
        """Sets up auth credentials to listen to twitter stream.

        :param config: TwitterConfig to use for authentication.
        :return: OAuthHandler
        """
        config = TwitterConfig()
        auth = OAuthHandler(config.CONSUMER_KEY, config.CONSUMER_SECRET)
        auth.set_access_token(config.ACCESS_TOKEN, config.ACCESS_TOKEN_SECRET)
        return auth

    def __get_tracked_stocks(self) -> Dict[str, TrackedStock]:
        """Get all tracked stocks as a dict with symbols as keys.

        :return: Tracked stocks as dict.
        """
        stock_repo = SQLStockRepo()
        stocks = stock_repo.get_all()
        return {
            s.symbol: TrackedStock(name=s.name, symbol=s.symbol) for s in stocks
        }

    def __setup_tweet_service(self) -> TweetService:
        """Sets up a usable tweet service.

        :return: TweetService
        """
        filter_svc = SpamFilterService(SpamFilterConfig())
        ranking_svc = MQRankingService(self.TRACKED_STOCKS, MQConfig())
        tweet_repo = SQLTweetRepo()
        return TweetServiceImpl(self.TRACKED_STOCKS, filter_svc, ranking_svc, tweet_repo)
