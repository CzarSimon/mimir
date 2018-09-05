# Standard library
import json
import logging
from abc import ABCMeta, abstractmethod
from urlparse import urlparse

# 3rd party modules
import requests

# Internal modules
from app.config import values


class RankingService(metaclass=ABCMeta):

    @abstractmethod
    def rank(self, tweet, links, symbols):
        """Sends tweet contents to be ranked.

        :param tweet: Tweet to rank.
        :param links: Links to extract and rank.
        :param symbols: Symbols to match against.
        """
        pass


class RankingServiceImpl(RankingService):

    __log = logging.getLogger('RankingServiceImpl')

    def __init__(self, tracked_stocks, config):
        self.TRACKED_STOCKS = tracked_stocks
        self.RANK_URL = f'{config.URL}{config.RANK_ROUTE}'
        self.HEADERS = {
            'Content-Type': 'application/json',
            'User-Agent': values.USER_AGENT
        }

    def rank(self, tweet, links, symbols):
        body = json.dumps(self.__create_rank_body(tweet, links, symbols))
        self.__log.info(body)
        resp = requests.post(self.RANK_URL, data=body, headers=self.HEADERS,
                             timeout=values.RPC_TIMEOUT)
        if not resp.ok:
            self.__log.error(f'Ranking failed: {resp.status} - {resp.text}')

    def __create_rank_body(tweet, links, symbols):
        """Creates a rank object.

        :param tweet: Tweet to rank.
        :param links: Links to extract and rank.
        :param symbols: Symbols to match against.
        :return: Rank object as a dict.
        """
        return {
            'urls': [link.url for link in links if self.__allowed_link(link)],
            'subjects': [self.TRACKED_STOCKS[s.symbol] for s in symbols],
            'author': {
                'id': tweet.author_id,
                'followerCount': tweet.author_followers
            },
            'language': tweet.language
        }

    def __allowed_link(self, link):
        """Checks if a tweet link points to an allowd domain.

        :param link: TweetLink to check.
        :return: Boolean indicating that the link is allowed.
        """
        return urlparse(link.url).netloc not in values.FORBIDDEN_DOMAINS
