# Standard library
import os
from logging.config import dictConfig

# Setup of logging configureaion
from .logging import LOGGING_CONIFG
dictConfig(LOGGING_CONIFG)

# Internal modultes
from app.config import util


CASHTAG_THRESHOLD = float(os.getenv('CASHTAG_THRESHOLD', '0.8'))
TRAIN_MODEL = os.getenv('TRAIN_MODEL', 'TRUE') == 'TRUE'


class AppConfig(object):
    SQLALCHEMY_DATABASE_URI = util.get_database_uri()
    SQLALCHEMY_TRACK_MODIFICATIONS = False
