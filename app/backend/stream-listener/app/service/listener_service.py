# Standard library
import logging
import json
import time
import sys
from threading import Thread

# 3rd party modules
from tweepy.streaming import StreamListener

# Internal modules
from app.config import values


class StreamListenerImpl(StreamListener):
    """Stream Listner that uses a tweet service for handling tweets."""

    __log = logging.getLogger('StreamListenerImpl')

    def __init__(self, tweet_svc, twitter_config):
        self.__tweet_svc = tweet_svc
        self.__error_count = 0
        self.RATE_LIMIT_CODE = twitter_config.RATE_LIMIT_CODE
        self.ERROR_PAUSE = twitter_config.ERROR_PAUSE_SECONDS

    def on_data(self, data):
        Thread(target=self.__tweet_svc.handle, args=(data,)).start()
        self.__error_count = 0

    def on_error(self, status_code):
        self.__log.error(f'Encountered error: {status_code}, exiting')
        pause_seconds = self.ERROR_PAUSE * self.__error_count
        if status_code == self.RATE_LIMIT_CODE:
            pause_seconds *= 2
        time.sleep(pause_seconds)


class StreamLogger(StreamListener):
    """Stream Listnener that logs incomming tweets to the configured log."""

    __log = logging.getLogger('StreamLogger')

    def on_data(self, data):
        formated_data = self.__format_data(data)
        self.__log.info(formated_data)

    def on_error(self, status_code):
        self.__log.error(f'Encountered error: {status_code}, exiting')
        sys.exit(1)

    def __format_data(self, data):
        deserialized_data = json.loads(data)
        return json.dumps(deserialized_data, indent=4, sort_keys=True)
