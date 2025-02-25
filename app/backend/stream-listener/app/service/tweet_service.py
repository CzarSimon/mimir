# Standard library
import json
import logging
from abc import ABCMeta, abstractmethod
from typing import Dict, List

# Internal modules
from app.config import values
from app.models import Tweet, TweetLink, TweetSymbol, TweetContent
from app.models import TrackedStock
from app.service import FilterService
from app.service import RankingService
from app.service import FilterService
from app.repository import TweetRepo


class TweetService(metaclass=ABCMeta):
    """Interface for handling incomming raw tweets."""

    @abstractmethod
    def handle(self, raw_tweet):
        """Handles parsing, filtering, storing and dispatching raw tweets.

        :param raw_tweet: Raw tweet dict to handle.
        """


class TweetServiceImpl(TweetService):

    __log = logging.getLogger('TweetServiceImpl')

    def __init__(self, tracked_symbols: Dict[str, TrackedStock],
                 filter_svc: FilterService, ranking_svc: RankingService,
                 tweet_repo: TweetRepo) -> None:
        self.TRACKED_SYMBOLS = tracked_symbols
        self.__filter_svc = filter_svc
        self.__ranking_svc = ranking_svc
        self.__tweet_repo = tweet_repo

    def handle(self, raw_tweet: bytes) -> None:
        content = self.__parse_tweet_contents(raw_tweet)
        if self.__filter_svc.is_spam(content.tweet):
            self.__log.info(f'SPAM: {content.tweet}')
            return
        self.__save_content(content)
        self.__ranking_svc.rank(content)

    def __parse_tweet_contents(self, raw_tweet) -> TweetContent:
        """Parses a raw tweet dict into a tweet, links and symbols.

        :param raw_tweet: Raw tweet to parse.
        :return: Parsed TweetContent
        """
        deserilized_tweet = json.loads(raw_tweet)
        tweet = self.__parse_tweet(deserilized_tweet)
        links = self.__parse_links(tweet.id, deserilized_tweet)
        symbols = self.__parse_symbols(tweet.id, deserilized_tweet)
        # assert len(symbols) != 0
        return TweetContent(tweet=tweet, links=links, symbols=symbols)

    def __parse_tweet(self, tweet_dict: Dict) -> Tweet:
        """Parses tweet as dict into the Tweet model structure.

        :param tweet_dict: Full tweet dictionary.
        :return: Parsed Tweet
        """
        return Tweet(text=tweet_dict['text'],
                     language=tweet_dict['lang'],
                     author_id=tweet_dict['user']['id_str'],
                     author_followers=tweet_dict['user']['followers_count'])

    def __parse_links(self, tweet_id: str, tweet_dict: Dict) -> List[TweetLink]:
        """Parses tweet as dict into a list of TweetLinks.

        :param tweet_id: Id of the parent tweet.
        :param tweet_dict: Full tweet dictionary.
        :return: List of TweetLinks
        """
        entities = tweet_dict['entities']
        urls = [self.__parse_url(url) for url in entities['urls']]
        full_urls = filter(lambda url: url != '' and url != None, urls)
        return [TweetLink(url=url, tweet_id=tweet_id) for url in full_urls]

    def __parse_url(self, url: Dict[str, str]) -> str:
        """Extracts url string from a dict of urls.

        :param url: URLs as a dict.
        :return: URL as a string.
        """
        if 'expanded_url' in url and url['expanded_url'] != None:
            return url['expanded_url']
        elif 'url' in url:
            return url['url']
        return ''

    def __parse_symbols(self, tweet_id: str, tweet: Dict) -> List[TweetSymbol]:
        """Parses tweet as dict into a list of TweetSymbol.

        :param tweet_id: Id of the parent tweet.
        :param tweet: Full tweet dictionary.
        :return: List of TweetSymbols
        """
        all_symbols = self.__parse_symbol_text(tweet)
        symbols = filter(lambda s: s in self.TRACKED_SYMBOLS, set(all_symbols))
        return [TweetSymbol(symbol=s, tweet_id=tweet_id) for s in symbols]

    def __parse_symbol_text(self, tweet: Dict) -> List[str]:
        """Parses symbols from a complete tweet.

        :param tweet: Raw tweet as dict.
        :return: List of stock symbols in tweet.
        """
        all_symbols = self.__parse_symbols_from_entities(tweet)
        if 'extended_tweet' in tweet:
            extended_tweet = tweet['extended_tweet']
            all_symbols += self.__parse_symbols_from_entities(extended_tweet)
        if 'retweeted_status' in tweet:
            all_symbols += self.__parse_symbol_text(tweet['retweeted_status'])
        return all_symbols

    def __parse_symbols_from_entities(self, component: Dict) -> List[str]:
        """Parses stock symbols from a tweet component.

        :param component: Tweet component as a dict to search through.
        :return: List of stock symbols as strings.
        """
        if 'entities' not in component:
            return []
        entities = component['entities']
        return [symbol['text'].upper() for symbol in entities['symbols']]

    def __save_content(self, tweet_content: TweetContent) -> None:
        """Stores tweet, links and symbols from a raw tweet.

        :param tweet_content: TweetContent to store.
        """
        self.__tweet_repo.save_tweet(tweet_content.tweet)
        self.__tweet_repo.save_links(tweet_content.links)
        self.__tweet_repo.save_symbols(tweet_content.symbols)
