# Standard library
import logging


class App(object):

    __log = logging.getLogger('App')

    def start(self):
        self.__log.info('Starting app')
