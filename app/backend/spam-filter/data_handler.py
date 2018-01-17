import config
import sys
import psycopg2
import re


def clean_training_data(text):
    words = text.split(' ')
    filtered_words = filter(lambda word: not (is_url(word) or is_cashtag(word)), words)
    return " ".join(filtered_words)


def clean_text(words):
    words_without_urls = filter(lambda word: not is_url(word), words)
    filtered_words, over_threshold = filter_cashtags(words_without_urls)
    return over_threshold, " ".join(filtered_words)


def is_cashtag(text):
    return text.startswith('$')


def is_url(text):
    pattern = "(https?|ftp):\/\/[\.[a-zA-Z0-9\/\-]+"
    return re.match(pattern, text) is not None


def filter_cashtags(words):
    words_without_cashtags = filter(lambda word: not is_cashtag(word), words)
    over_threshold = is_over_cashtag_threshold(len(words), len(words_without_cashtags))
    return words_without_cashtags, over_threshold


def is_over_cashtag_threshold(start_length, end_length):
    cashtag_ratio = 1.0 - (float(end_length) / float(start_length))
    return cashtag_ratio > config.cashtag_threshold


def structure_training_data(raw_data):
    return dict(
        data=map(lambda datapoint: clean_training_data(datapoint[0]), raw_data),
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
    database = config.database
    conn_string = "host='{}' dbname='{}' user='{}' password='{}' port='{}'".format(
        database["HOST"], database["NAME"], database["USER"], database["PASSWORD"], database["PORT"])
    try:
        conn = psycopg2.connect(conn_string)
    except Exception as e:
        print e.message
        sys.exit(1)
    return conn
