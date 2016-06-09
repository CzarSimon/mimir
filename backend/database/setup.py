def _setupTables(confirmed):
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

if __name__ == "__main__":
    print "Type 1 to setup tables"
    should_setup = input()
    confirmation = True if (should_setup == 1) else False
    _setupTables(confirmed=confirmation)