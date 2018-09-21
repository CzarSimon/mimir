# 3rd party modules.
from flasgger import swag_from
from flask import jsonify, make_response

# Internal modules
from app import app
from app import controllers
from app.controllers import classification, training_data
from app.controllers import status, errors


@app.route('/v1/classify', methods=['POST'])
@swag_from('swagger/v1-classify.yml')
def classify_spam():
    result = classification.is_spam()
    return _create_response(result)


@app.route('/v1/training-data', methods=['POST'])
@swag_from('swagger/v1-training-data.yml')
def add_training_data():
    training_data.add_training_data()
    return _create_ok_response()


@app.route('/health', methods=['GET'])
def check_health():
    result, status = controllers.health_check.check_health()
    return _create_response(result, status)


def _create_response(result, status=status.HTTP_200_OK):
    """Returns a response indicating that an index update was triggered.

    :return: flask.Response.
    """
    return make_response(jsonify(result), status)


def _create_ok_response():
    """Creates a 200 OK response.

    :return: flask.Response.
    """
    return make_response(jsonify({'status': 'OK'}), status.HTTP_200_OK)
