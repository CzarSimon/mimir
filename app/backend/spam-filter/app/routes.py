# 3rd party modules.
from flask import jsonify, make_response

# Internal modules
from app import app
from app import controllers
from app.controllers import status, errors


@app.route('/v1/classify', methods=['POST'])
def classify_spam():
    raise errors.NotImplementedError()


@app.route('/v1/training-data', methods=['POST'])
def add_training_data():
    raise errors.NotImplementedError()


@app.route('/health', methods=['GET'])
def check_health():
    result, status = controllers.health_check.check_health()
    return _create_response(result, status)


def _create_response(result, status=status.HTTP_200_OK):
    """Returns a response indicating that an index update was triggered.

    :return: flask.Response.
    """
    return make_response(jsonify(result), status)
