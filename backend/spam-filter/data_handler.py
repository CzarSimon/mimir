import config
import sys
import psycopg2
import re


def clean_text(words):
    filtered_words = filter(lambda word: not is_url(word) and not is_cashtag(word), words)
    return " ".join(filtered_words)


def is_cashtag(text):
    return text.startswith('$')


def is_url(text):
    pattern = "(https?|ftp):\/\/[\.[a-zA-Z0-9\/\-]+"
    return re.match(pattern, text) is not None


def structure_training_data(raw_data):
    return dict(
        data=map(lambda datapoint: clean_text(datapoint[0].split(" ")), raw_data),
        labels=map(lambda datapoint: datapoint[1], raw_data)
    )


def _encode(datapoint):
    return datapoint["text"].encode('utf-8')


def get_training_data():
    conn = _connect_pg()
    data = None
    try:
        cursor = conn.cursor()
        cursor.execute("SELECT TWEET, LABEL FROM SPAM_DATA")
        data = cursor.fetchall()
    except Exception as e:
        print e.message
        conn.rollback()
        sys.exit(1)
    finally:
        cursor.close()
        conn.close()
        return structure_training_data(data)


def _connect_pg():
    conn = None
    database = config.postgres
    conn_string = "host='{}' dbname='{}' user='{}' password='{}'".format(database["HOST"], database["NAME"], database["USER"], database["PASSWORD"])
    try:
        conn = psycopg2.connect(conn_string)
    except Exception as e:
        print e.message
        sys.exit(1)
    return conn
