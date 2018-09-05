# Standard library
from abc import ABCMeta, abstractmethod

# Internal modules
from app import db


class TweetRepo(metaclass=ABCMeta):
    """Interface for storage and retrieval of Tweet entinties."""

    @abstractmethod
    def save_tweet(self, tweet):
        """Stores a tweet.

        :param tweet: Tweet to store.
        """
        pass

    @abstractmethod
    def save_links(self, links):
        """Stores the links in a tweet.

        :param links: List of TweetLinks to store.
        """
        pass

    @abstractmethod
    def save_symbols(self, symbols):
        """Stores symbols referenced in a tweet.

        :param symbols: List of TweetSymbols to store.
        """
        pass


class SQLTweetRepo(TweetRepo):
    """TweetRepo implemented against a SQL database."""

    def save_tweet(self, tweet):
        """Stores a tweet.

        :param tweet: Tweet to store.
        """
        db.session.add(tweet)
        db.session.commit()

    def save_links(self, links):
        """Stores the links in a tweet.

        :param links: List of TweetLinks to store.
        """
        for link in links:
            db.session.add(link)
        db.session.commit()

    def save_symbols(self, symbols):
        """Stores symbols referenced in a tweet.

        :param symbols: List of TweetSymbols to store.
        """
        for symbol in symbols:
            db.session.add(symbol)
        db.session.commit()
