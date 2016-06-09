import sys
import psycopg2
from datetime import datetime
from datetime import date
from credentials import database
sys.path.append("..")

conn_string = "host='{}' dbname='{}' user='{}' password='{}'".format(database["HOST"], database["NAME"], database["USER"], database["PASSWORD"])
print conn_string
conn = psycopg2.connect(conn_string)

def closeDatabase():
    conn.close()

def queryDatabase(query, multiple):
    return_item = True
    try:
        c = conn.cursor()
        c.execute(query)
        if multiple:
            return_item = c.fetchall()
        else:
            return_item = c.fetchone()
    except Exception as e:
        print 'queryDatabase in databaseManger failed. ', e
        return_item = False
    finally:
        c.close()
        return return_item

def insertTweet(tweetId, userId, date, tweet, ticker, urls, lang, followers):
    stockTweet = (tweetId, userId, date, tweet, ticker, urls, lang, followers)
    c = conn.cursor()
    try:
        c.execute('INSERT INTO stockTweets (tweetId, userId, createdAt, tweet, ticker, urls, lang, followers) VALUES (%s, %s, %s, %s, %s, %s, %s, %s)', stockTweet)
        conn.commit()
    except Exception as e:
        print "-#-#-#-#-#-#-#-the problem occured here"
        print e
        conn.rollback()
    finally:
        c.close()
        return True

def updateMeanAndStdev(tic, newMean, newStdev):
    c = conn.cursor()
    success = True
    try:
        c.execute('UPDATE stocks SET mean=%s, stdev=%s WHERE ticker=%s', (newMean, newStdev, tic))
        conn.commit()
    except Exception as e:
        print e
        conn.rollback()
        success = False
    finally:
        return success

def addAlias(alias, ticker):
    c = conn.cursor()
    newAlias = (alias, ticker)
    success = True
    try:
        c.execute('INSERT INTO tickerAliases VALUES (%s, %s)', newAlias)
        conn.commit()
    except Exception as e:
        print e
        success = False
    finally:
        c.close()
        return success

def logActiveDate():
    print "Logging date"
    c = conn.cursor()
    success = True
    curr_date = datetime.utcnow()
    format_date = date(curr_date.year, curr_date.month, curr_date.day)
    try:
        c.execute('INSERT INTO activeDates VALUES (%s)', (format_date,))
        conn.commit()
    except Exception as e:
        print e
        conn.rollback()
        success = False
        print "no success"
    finally:
        c.close()
        return success
