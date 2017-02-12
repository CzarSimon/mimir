import psycopg2, sys
from config import database

conn_string = "host='{}' dbname='{}' user='{}' password='{}'".format(database["HOST"], database["NAME"], database["USER"], database["PASSWORD"])
try:
    conn = psycopg2.connect(conn_string)
except Exception as e:
    print e
    sys.exit(1)

def get_hourly_volume(datetime):
    query = "SELECT ticker, COUNT(*) FROM stocktweets WHERE createdAt>'{}' GROUP BY ticker".format(datetime)
    return _execute_query(query, "Hourly volume count failed in db retrival")

def get_unique_tickers():
    query = "SELECT ticker FROM stocks WHERE is_tracked=TRUE"
    return _execute_query(query, "Intital retrival of unique stock tickers failed")

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
