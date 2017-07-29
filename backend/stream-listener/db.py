# standard library
import sys
from datetime import datetime
from datetime import date

# external library
import psycopg2

# user specific inport
from config import database


conn_string = "host='{}' dbname='{}' user='{}' password='{}'".format(database["HOST"], database["NAME"], database["USER"], database["PASSWORD"])


try:
    print conn_string
    conn = psycopg2.connect(conn_string)
except Exception as e:
    print "Db connedction failed: " + str(e)
    sys.exit(1)


# close_database Closes a database connection
def close_database():
    conn.close()


# query_db Executes a query against a database and returns the result(s)
def query_db(query, multiple = True):
    return_item = True
    try:
        c = conn.cursor()
        c.execute(query)
        if multiple:
            return_item = c.fetchall()
        else:
            return_item = c.fetchone()
    except Exception as e:
        print 'query_db in databaseManger failed. ', e
        return_item = False
    finally:
        c.close()
        return return_item


# insert_tweet Inserts a tweet in the database
def insert_tweet(tweetId, userId, date, tweet, ticker, lang, followers):
    stockTweet = (tweetId, userId, date, tweet, ticker, lang, followers)
    c = conn.cursor()
    try:
        c.execute('INSERT INTO stockTweets (tweetId, userId, createdAt, tweet, ticker, lang, followers) VALUES (%s, %s, %s, %s, %s, %s, %s)', stockTweet)
        conn.commit()
    except Exception as e:
        print "-#-#-#-#-#-#-#-the problem occured here"
        print e
        conn.rollback()
    finally:
        c.close()
        return True


# insert_untracked Records the occurence of an untracked tweet
def insert_untracked(tweet_id, ticker, timestamp):
    untracked_record = (tweet_id, ticker, timestamp)
    c = conn.cursor()
    try:
        c.execute("INSERT INTO untracked_tickers (id, ticker, timestamp) VALUES (%s, %s, %s)", untracked_record)
        conn.commit()
    except Exception as e:
        print(e)
        conn.rollback()
    finally:
        c.close()
