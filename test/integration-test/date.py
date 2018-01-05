from datetime import datetime


def utcnow():
    return datetime.utcnow()


def date_str(date):
    return date.strftime("%Y-%m-%dT%H:%M:%SZ")


def utc_str():
    return date_str(utcnow())
