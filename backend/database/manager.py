import sys
import psycopg2
from datetime import datetime
from datetime import date
from pytz import timezone
from credentials import database
sys.path.append("..")

conn_string = "host='{}' dbname='{}' user='{}' password='{}'".format(database["HOST"], database["NAME"], database["USER"], database["PASSWORD"])
print conn_string
conn = psycopg2.connect(conn_string)

def closeDatabase():
    conn.close()

def setupTables(confirmed):
    if confirmed:
        c = conn.cursor()

        c.execute('DROP TABLE IF EXISTS activeDates')
        c.execute('DROP TABLE IF EXISTS tickerAliases')
        c.execute('DROP TABLE IF EXISTS stockTweets')
        c.execute('DROP TABLE IF EXISTS stocks')

        c.execute('CREATE TABLE stocks (ticker TEXT PRIMARY KEY, name TEXT, storedAt DATE, mean NUMERIC(7,2)[2][24], stdev NUMERIC(7,2)[2][24])')
        empty_list = [0.0] * 24
        mean = [empty_list, empty_list]
        stdev = mean
        current_date = datetime.now(tz=timezone('UTC'))
        initial_stocks = [
            ('$LNKD', 'LinkedIn Corporation', current_date, mean, stdev),
            ('$FB', 'Facebook Inc.', current_date, mean, stdev),
            ('$TWTR', 'Twitter, Inc.', current_date, mean, stdev),
            ('$ACN', 'Accenture plc', current_date, mean, stdev),
            ('$AAPL', 'Apple Inc.', current_date, mean, stdev),
            ('$NKE', 'Nike, Inc.', current_date, mean, stdev),
            ('$AMZN', 'Amazon.com, Inc.', current_date, mean, stdev),
            ('$NFLX', 'Netflix, Inc.', current_date, mean, stdev),
            ('$MSFT', 'Microsoft Corporation', current_date, mean, stdev),
            ('$TSLA', 'Tesla Motors, Inc.', current_date, mean, stdev),
            ('$INTC', 'Intel Corporation', current_date, mean, stdev),
            ('$WMT', 'Wal-Mart Stores Inc.', current_date, mean, stdev),
            ('$GS', 'The Goldman Sachs Group, Inc.', current_date, mean, stdev),
            ('$SCTY', 'SolarCity Corporation', current_date, mean, stdev),
            ('$T', 'AT&T, Inc.', current_date, mean, stdev),
            ('$YELP', 'Yelp Inc.', current_date, mean, stdev),
            ('$EBAY', 'eBay Inc.', current_date, mean, stdev),
            ('$PYPL', 'PayPal Holdings, Inc.', current_date, mean, stdev),
            ('$GOOG', 'Alphabet Inc.', current_date, mean, stdev)
        ]
        record_list_template = ','.join(['%s'] * len(initial_stocks))
        insert_query = 'INSERT INTO stocks (ticker, name, storedAt, mean, stdev) VALUES {0}'.format(record_list_template)
        c.execute(insert_query, initial_stocks)

        c.execute('CREATE TABLE stockTweets (tweetId TEXT, userId TEXT, createdAt TIMESTAMP, tweet TEXT, ticker TEXT REFERENCES stocks(ticker), urls TEXT, lang TEXT, followers INTEGER, PRIMARY KEY (tweetId, ticker))')
        c.execute('CREATE TABLE tickerAliases (alias text PRIMARY KEY, ticker text REFERENCES stocks(ticker))')
        c.execute('CREATE TABLE activeDates (activeDate DATE PRIMARY KEY)')
        c.execute("INSERT INTO tickerAliases VALUES ('$GOOGL', '$GOOG')")

        conn.commit()
        c.close()
        print "Changes made"
    else:
        print "No changes made"

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
