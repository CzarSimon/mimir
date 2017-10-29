# external library
import rethinkdb as r
import psycopg2 as pg

# standard library
from datetime import datetime
from os import getenv
import json
import time


def pg_connect():
    conn_string = build_pg_conn_string()
    return pg.connect(conn_string)


def build_pg_conn_string():
    host = getenv("PG_HOST", "localhost")
    db_name = getenv("PG_NAME", "mimirprod")
    user = getenv("PG_USER", "simon")
    password = getenv("PG_PASSWORD", "pwd")
    return "host='{}' dbname='{}' user='{}' password='{}'".format(
    host, db_name, user, password)


def rdb_connect():
    host = getenv("RDB_HOST", "localhost")
    db_name = getenv("RDB_NAME", "mimir_app_server")
    return r.connect(host=host, db=db_name)


def migrate_stocks(rdb_conn, pg_conn):
    query = "INSERT INTO STOCK (TICKER, NAME, DESCRIPTION, IMAGE_URL, WEBSITE) VALUES (%s, %s, %s, %s, %s)"
    cur = pg_conn.cursor()
    success = False
    for stock in r.table("stocks").run(rdb_conn):
        values = get_stock_values(stock)
        success = insert_row(query, values, cur, pg_conn)
        if not success:
            break
    cur.close()
    if success:
        pg_conn.commit()


def get_stock_values(stock):
    return (
        stock["ticker"],
        stock["name"],
        get_value(stock, "description"),
        get_value(stock, "imageUrl"),
        get_value(stock, "website")
    )


def migrate_twitter_data(rdb_conn, pg_conn):
    query = '''
    INSERT INTO TWITTER_DATA (
        TICKER, MINUTE, VOLUME, BUSDAY_MEAN, WEEKEND_MEAN, BUSDAY_STDEV, WEEKEND_STDEV
    ) VALUES (%s, %s, %s, %s, %s, %s, %s)'''
    cur = pg_conn.cursor()
    success = False
    for stock in r.table("stocks").run(rdb_conn):
        values = get_twitter_data_values(stock)
        success = insert_row(query, values, cur, pg_conn)
        if not success:
            break
    cur.close()
    if success:
        pg_conn.commit()


def get_twitter_data_values(stock):
    return (
        stock["ticker"],
        get_value(stock, "minute"),
        get_value(stock, "volume"),
        get_value(stock["mean"], "busdays"),
        get_value(stock["mean"], "weekend_days"),
        get_value(stock["mean"], "busdays"),
        get_value(stock["mean"], "weekend_days")
    )


def migrate_users(rdb_conn, pg_conn):
    query = "INSERT INTO APP_USER (ID, EMAIL, TICKERS, JOIN_DATE) VALUES (%s, %s, %s, %s)"
    cur = pg_conn.cursor()
    success = False
    for user in r.table("users").run(rdb_conn):
        values = get_user_values(user)
        success = insert_row(query, values, cur, pg_conn)
        success = migrate_sessions(user, cur, pg_conn)
        success = migrate_search_history(user, cur, pg_conn)
        if not success:
            break
    cur.close()
    if success:
        pg_conn.commit()


def get_user_values(user):
    return (
        user["id"],
        get_value(user, "email"),
        get_value(user, "tickers"),
        user["joinDate"]
    )


def migrate_sessions(user, cursor, conn):
    query = "INSERT INTO SESSION (USER_ID, SESSION_START) VALUES (%s, %s)"
    user_id = user["id"]
    success = True
    for session_start in get_value(user, "sessions"):
        success = insert_row(query, (user_id, session_start), cursor, conn)
        if not success:
            break
    return success


def migrate_search_history(user, cursor, conn):
    query = "INSERT INTO SEARCH_HISTORY (USER_ID, SEARCH_TERM, DATE_INSERTED) VALUES (%s, %s, %s)"
    user_id = user["id"]
    success = True
    now = datetime.utcnow()
    for search_term in get_value(user, "searchHistory"):
        success = insert_row(query, (user_id, search_term, now), cursor, conn)
        if not success:
            break
    return success


def get_value(dictonary, key):
    return dictonary[key] if key in dictonary else None


def insert_row(query, values, cursor, conn):
    success = True
    try:
        cursor.execute(query, values)
    except Exception as e:
        print e
        success = False
        conn.callback()
    finally:
        return success


def migrate():
    rdb_conn = rdb_connect()
    pg_conn = pg_connect()
    migrate_stocks(rdb_conn, pg_conn)
    migrate_twitter_data(rdb_conn, pg_conn)
    migrate_users(rdb_conn, pg_conn)
    rdb_conn.close()
    pg_conn.close()


def main():
    migrate()
    while True:
        print "Done, sleeping"
        time.sleep(10)

if __name__ == '__main__':
    main()
