# Standard library
import json
import logging
from abc import ABCMeta, abstractmethod

# 3rd party modules
import requests

# Internal modules
from app.config import values


class FilterService(metaclass=ABCMeta):

    @abstractmethod
    def is_spam(self, tweet):
        """Sends tweet contents to be ranked.

        :param tweet: Tweet to check.
        :return: Boolean indicating if tweet is spam.
        """
        pass


class SpamFilterService(FilterService):

    __log = logging.getLogger('SpamFilterService')

    def __init__(self, config):
        self.CLASSIFY_URL = f'{config.URL}{config.CLASSIFY_ROUTE}'
        self.HEADERS = {
            'Content-Type': 'application/json',
            'User-Agent': values.USER_AGENT
        }

    def is_spam(self, tweet):
        spam_candidate = self.__create_spam_body(tweet)
        self.__log.info(spam_candidate)
        body = json.dumps(spam_candidate)
        self.__log.info(body)
        resp = requests.post(self.CLASSIFY_URL, data=body,
                             headers=self.HEADERS, timeout=values.RPC_TIMEOUT)
        if not resp.ok:
            self.__log.error(f'Ranking failed: {resp.status} - {resp.text}')
            return False
        return self.__tweet_was_spam(resp)

    def __create_spam_body(self, tweet):
        """Formats a tweet into a spam candidate.

        :param tweet: Tweet to check.
        :return: Boolean indicating
        """
        return {
            'text': tweet.text
        }

    def __tweet_was_spam(self, resp):
        """Parses spam filter response to check if spam was detected.

        :param resp: Spam filter response.
        :return: Boolean indicating if tweet is spam.
        """
        try:
            resp_body = resp.json()
            return resp_body['label'] == values.SPAM_LABEL
        except Exception as e:
            self.__log.error(str(e))
            return False
