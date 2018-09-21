# Standard library
import logging

# 3rd party modules.
from flask import Flask, jsonify, make_response, request
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate
from flasgger import Swagger

# Internal modules
from app.config import AppConfig

app = Flask('Mimir Spam Filter')
app.config.from_object(AppConfig)
db = SQLAlchemy(app)
migrate = Migrate(app, db)
swagger = Swagger(app)


from app import routes, models
from app.controllers import errors


_error_log = logging.getLogger('ErrorHandler')


@app.errorhandler(errors.RequestError)
def handle_request_error(error):
    """Handles errors encountered when handling requests.

    :param error: Encountered RequestError.
    :return: flask.Response indicating the encountered error.
    """
    _error_log.warning(str(error))
    json_error = jsonify(error=str(error),
                         status=error.status(),
                         path=request.path)
    return make_response(json_error, error.status())
