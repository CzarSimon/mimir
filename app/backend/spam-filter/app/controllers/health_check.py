# Standard library
import logging

# Internal modules
from app import db
from app.controllers import status


_log = logging.getLogger(__name__)


STATUS_UP = 'UP'
STATUS_DOWN = 'DOWN'


def check_health():
    """Handles health check requests.

    :return: Status info as a dict.
    :return: HTTP status code.
    """
    model_status, model_ok = _check_model_status()
    overall_status, status_code = _determine_overall_status(model_ok)
    return {'status': overall_status, 'model': model_status}, status_code


def _check_model_status():
    """Checks that the spam filter model is trained
    and ready to classify requests

    :return: Status text.
    :return: Boolean indication if status is healthy
    """
    return STATUS_DOWN, False


def _determine_overall_status(*statuses):
    """Checks that all statuses entered are true for healthy.

    :param statuses: List of booelans indicating individual components health.
    :return: Status description as a string
    :return: Status code.
    """
    if all(statuses):
        return STATUS_UP, status.HTTP_200_OK
    return STATUS_DOWN, status.HTTP_503_SERVICE_UNAVAILIBLE
