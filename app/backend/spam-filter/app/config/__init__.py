# Standard library
import os

# Internal modultes
from app.config import util


CASHTAG_THRESHOLD = float(os.getenv('CASHTAG_THRESHOLD', '0.8'))


class AppConfig(object):
    SQLALCHEMY_DATABASE_URI = util.get_database_uri()
    SQLALCHEMY_TRACK_MODIFICATIONS = False
