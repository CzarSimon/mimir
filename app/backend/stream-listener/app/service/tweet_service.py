# Standard library
import json
import logging
from abc import ABCMeta, abstractmethod

# Internal modules
from app.config import values
from app.models import Tweet, TweetLink, TweetSymbol


class TweetService(metaclass=ABCMeta):
    """Interface for handling incomming raw tweets."""

    @abstractmethod
    def handle(self, raw_tweet):
        """Handles parsing, filtering, storing and dispatching raw tweets.

        :param raw_tweet: Raw tweet dict to handle.
        """
        pass


class TweetServiceImpl(TweetService):

    __log = logging.getLogger('TweetServiceImpl')

    def __init__(self, tracked_symbols, filter_svc, ranking_svc, tweet_repo):
        self.TRACKED_SYMBOLS = tracked_symbols
        self.__filter_svc = filter_svc
        self.__ranking_svc = ranking_svc
        self.__tweet_repo = tweet_repo

    def handle(self, raw_tweet):
        tweet, links, symbols = self.__parse_tweet_contents(raw_tweet)
        self.__log.info(f'{tweet}')
        self.__log.info(f'{links}')
        self.__log.info(f'{symbols}')
        if self.__filter_svc.is_spam(tweet):
            self.__log.info(f'SPAM: {tweet}')
            return
        self.__save_content(tweet, links, symbols)
        self.__ranking_svc.rank(tweet, links, symbols)

    def __parse_tweet_contents(self, raw_tweet):
        """Parses a raw tweet dict into a tweet, links and symbols.

        :param raw_tweet: Raw tweet to parse.
        :return: Parsed Tweet
        :return: Parsed list of TweetLinks
        :return: Parsed list of TweetSymbols
        """
        deserilized_tweet = json.loads(raw_tweet)
        tweet = self.__parse_tweet(deserilized_tweet)
        links = self.__parse_links(tweet.id, deserilized_tweet)
        symbols = self.__parse_symbols(tweet.id, deserilized_tweet)
        return tweet, links, symbols

    def __parse_tweet(self, tweet_dict):
        """Parses tweet as dict into the Tweet model structure.

        :param tweet_dict: Full tweet dictionary.
        :return: Parsed Tweet
        """
        return Tweet(text=tweet_dict['text'],
                     language=tweet_dict['user']['id_str'],
                     author_id=tweet_dict['user']['followers_count'],
                     author_followers=tweet_dict['lang'])

    def __parse_links(self, tweet_id, tweet_dict):
        """Parses tweet as dict into a list of TweetLinks.

        :param tweet_id: Id of the parent tweet.
        :param tweet_dict: Full tweet dictionary.
        :return: List of TweetLinks
        """
        entities = tweet_dict['entities']
        urls = [self.__parse_url(url) for url in entities['urls']]
        full_urls = filter(lambda url: url != '' and url != None, urls)
        return [TweetLink(url=url, tweet_id=tweet_id) for url in full_urls]

    def __parse_url(self, url):
        """Extracts url string from a dict of urls.

        :param url: URLs as a dict.
        :return: URL as a string.
        """
        if 'expanded_url' in url and url['expanded_url'] != None:
            return url['expanded_url']
        elif 'url' in url:
            return url['url']
        return ''

    def __parse_symbols(self, tweet_id, tweet_dict):
        """Parses tweet as dict into a list of TweetSymbol.

        :param tweet_id: Id of the parent tweet.
        :param tweet_dict: Full tweet dictionary.
        :return: List of TweetSymbols
        """
        entities = tweet_dict['entities']
        all_symbols = [s['text'].upper() for s in entities['symbols']]
        symbols = filter(lambda s: s in self.TRACKED_SYMBOLS, all_symbols)
        return [TweetSymbol(symbol=s, tweet_id=tweet_id) for s in symbols]

    def __store_content(self, tweet, links, symbols):
        """Stores tweet, links and symbols from a raw tweet.

        :param tweet: Tweet to store.
        :param links: List of TweetLinks to store.
        :param symbols: List of TweetSymbols to store.
        """
        self.__tweet_repo.save_tweet(tweet)
        self.__tweet_repo.save_links(links)
        self.__tweet_repo.save_symbols(symbols)
