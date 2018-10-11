# Standard library
import json
import logging
from abc import ABCMeta, abstractmethod
from urllib.parse import urlparse
from typing import Dict, List

# 3rd party modules
import requests

# Internal modules
from app.config import values, NewsRankerConfig, MQConfig
from app.models import TweetContent, Tweet, TweetLink, TweetSymbol
from app.models import TrackedStocks


class RankingService(metaclass=ABCMeta):

    @abstractmethod
    def rank(self, tweet_content: TweetContent) -> None:
        """Sends tweet contents to be ranked.

        :param tweet_content: TweetContent.
        """

class MQRankingService(RankingService):

    __log = logging.getLogger('MQRankingService')

    def __init__(self, tracked: TrackedStocks, config: MQConfig) -> None:
        self.TRACKED_STOCKS = tracked

    def rank(self, tweet_content: TweetContent) -> None:
        pass

    def __create_rank_body(self, content: TweetContent) -> bytes:
        subjects = [self.TRACKED_STOCKS[s.symbol] for s in content.symbols]
        return json.dumps(create_rank_body(content, subjects))


class RestRankingService(RankingService):

    __log = logging.getLogger('RestRankingService')

    def __init__(self, tracked: TrackedStocks, config: NewsRankerConfig) -> None:
        self.TRACKED_STOCKS = tracked
        self.RANK_URL = f'{config.URL}{config.RANK_ROUTE}'
        self.HEADERS = {
            'Content-Type': 'application/json',
            'User-Agent': values.USER_AGENT
        }

    def rank(self, tweet_content: TweetContent) -> None:
        body = self.__create_rank_body(tweet, links, symbols)
        resp = requests.post(self.RANK_URL, data=body, headers=self.HEADERS,
                             timeout=values.RPC_TIMEOUT)
        if not resp.ok:
            self.__log.error(f'Ranking failed: {resp.status} - {resp.text}')

    def __create_rank_body(self, content: TweetContent) -> bytes:
        subjects = [self.TRACKED_STOCKS[s.symbol] for s in content.symbols]
        return json.dumps(create_rank_body(tweet, links, subjects))


def create_rank_body(content: TweetContent, subjects: List[TrackedStock]) -> Dict:
    """Creates a rank object.

    :param content: TweetContent.
    :param subjects: TrackedStocks to match against.
    :return: Rank object as a dict.
    """
    tweet = content.tweet
    return {
        'urls': [link.url for link in content.links if allowed_link(link)],
        'subjects': [sub.asdict() for sub in subjects],
        'author': {
            'id': tweet.author_id,
            'followerCount': tweet.author_followers
        },
        'language': tweet.language
    }

def allowed_link(link: str) -> Boolean:
    """Checks if a tweet link points to an allowd domain.

    :param link: TweetLink to check.
    :return: Boolean indicating that the link is allowed.
    """
    return urlparse(link.url).netloc not in values.FORBIDDEN_DOMAINS
