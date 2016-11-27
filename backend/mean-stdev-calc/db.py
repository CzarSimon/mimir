import psycopg2, sys
from config import database

conn_string = "host='{}' dbname='{}' user='{}' password='{}'".format(database["HOST"], database["NAME"], database["USER"], database["PASSWORD"])
try:
    conn = psycopg2.connect(conn_string)
except Exception as e:
    print e
    sys.exit(1)

def get_unique_tickers():
    query = "select ticker from stocks"
    return _execute_query(query, "Intital retrival of unique stock tickers failed")

def get_tweets_for_ticker(ticker):
    query = "select createdAt from stocktweets where ticker='{}'".format(ticker);
    return _execute_query(query, "Retrival of tweets for {} failed".format(ticker))

def get_first_day_stored(ticker):
    query = "select storedAt from stocks where ticker='{}'".format(ticker)
    return _execute_query(query, "Retrival of inital store date for {} failed".format(ticker))

def close_connection():
    conn.close();

def _execute_query(query, exeption_message):
    result = { "success": True, "data": None }
    try:
        cur = conn.cursor()
        cur.execute(query)
        result["data"] = cur.fetchall()
        conn.commit()
    except Exception as e:
        print exeption_message, e
        result["data"] = exeption_message
        result["success"] = False
        conn.rollback()
    finally:
        cur.close()
        return result
