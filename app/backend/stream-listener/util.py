# standard library import
from datetime import datetime
import json


# utcnow Returns the current timestamp in utc timezone
def utcnow():
    return str(datetime.utcnow())


# pretty_print Prints a dictonary as formated json
def pretty_print(dictonary):
    print json.dumps(dictonary, indent=4, sort_keys=True)


def is_empty(entity):
    return entity is None or entity == ""
