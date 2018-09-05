# Standard library
import logging
import sys

# Internal modules
from app import stream_listner


_log = logging.getLogger(__name__)


def main():
    try:
        stream_listner.start()
    except Exception as e:
        _log.info('Exiting: {}'.format(str(e)))
        sys.exit(1)


if __name__ == '__main__':
    main()
