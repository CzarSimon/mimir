import sys
sys.path.append("..")
from database import manager as db
import threading
from datetime import date
import numpy
import math, time

def getStockTweets():
    stockMeans = {}
    stockStdevs = {}
    threads = {}
    stockList = db.queryDatabase('SELECT ticker, createdAt FROM stockTweets ORDER BY ticker', True)
    daysStored = daysMeasured()
    tickerList = []
    prevTicker = stockList[0][0]
    for stock in stockList:
        if stock[0] == prevTicker:
            created_at = str(stock[1].year)+"-"+str(stock[1].month)+"-"+str(stock[1].day)+":"+str(stock[1].hour)
            tickerList += [created_at]
        else:
            days = daysStored[prevTicker]
            threads[prevTicker] = dispatch(tickerList, prevTicker, days, stockMeans, stockStdevs)
            prevTicker = stock[0]
            tickerList = []
    days = daysStored[prevTicker]
    threads[prevTicker] = dispatch(tickerList, prevTicker, days, stockMeans, stockStdevs)
    threadsFinished = False
    while not threadsFinished:
        threadsFinished = True
        for ticker, thread in threads.iteritems():
            if thread.is_alive():
                threadsFinished = False
    updateDB(stockMeans, stockStdevs)
    db.logActiveDate()
    return True

def dispatch(list, ticker, days, mean, stdev):
    d = threading.Thread(name=str(ticker) + " thread", target=reduceByHour, args=(list, ticker, days, mean, stdev))
    d.setDaemon(True)
    d.start()
    return d

def reduceByHour(list, ticker, days, meanResult, stdevResult):
    mentionsPerHourEachDay = {}
    for dateHour in list:
        if dateHour in mentionsPerHourEachDay:
            mentionsPerHourEachDay[dateHour] += 1
        else:
            mentionsPerHourEachDay[dateHour] = 1
    result = splitBusAndWeekendDays(mentionsPerHourEachDay, days)
    meanResult[ticker] = result["mean"]
    stdevResult[ticker] = result["stdev"]

def splitBusAndWeekendDays(dict, days):
    mentions_per_hour_each_busday = {}
    mentions_per_hour_each_weekend_day = {}
    for dateHour, count in dict.iteritems():
        this_date = map(int, dateHour.split(":")[0].split("-"))
        if numpy.is_busday([date(this_date[0], this_date[1], this_date[2])])[0]:
            mentions_per_hour_each_busday[dateHour] = count
        else:
            mentions_per_hour_each_weekend_day[dateHour] = count
    bday_res = reduceByDay(mentions_per_hour_each_busday, days["busdays"])
    wday_res = reduceByDay(mentions_per_hour_each_weekend_day, days["weekend_days"])
    return {"mean": [bday_res["mean"], wday_res["mean"]], "stdev": [bday_res["stdev"], wday_res["stdev"]]}

def reduceByDay(dict, days):
    hourlyMentions = {}
    for date, count in dict.iteritems():
        hour = int(date.split(":")[1])
        if hour in hourlyMentions:
            hourlyMentions[hour] += [float(count)]
        else:
            hourlyMentions[hour] = [float(count)]
    for h in range(0,24):
        if h in hourlyMentions:
            length = len(hourlyMentions[h])
            if length < days:
                hourlyMentions[h] += [0.0]*(days - length)
        else:
            hourlyMentions[h] = [0.0]*days
    return calcMeanAndStdev(hourlyMentions, days)

def calcMeanAndStdev(hourlyMentions, days):
    meanList = [0.0] * 24
    stdevList = [0.0] * 24
    if (days > 1):
        f_days = float(days)
        for hour, volumes in hourlyMentions.iteritems():
            res = calc(volumes, f_days)
            meanList[hour] = res["mean"]
            stdevList[hour] = res["stdev"]
    return {"mean": meanList, "stdev": stdevList}

def calc(seq, n):
    mean = sum(seq) / n
    stdev = sum(list(map(lambda x: (x - mean) ** 2, seq))) / (n - 1.0)
    return {"mean": round(mean, 2), "stdev": round(stdev, 2)}

def daysMeasured():
    endDate = datetimeToDate(db.queryDatabase('SELECT MAX(createdAt) FROM stockTweets', False)[0])
    startDate = datetimeToDate(db.queryDatabase('SELECT MIN(createdAt) FROM stockTweets', False)[0])
    bus_days = int(numpy.busday_count(startDate, endDate)) + 1
    all_days = (endDate - startDate).days + 1
    days = { "busdays": bus_days, "weekend_days": (all_days - bus_days) }
    stocks = db.queryDatabase('SELECT ticker, storedAt FROM stocks', True)
    inDb = {}
    for stock in stocks:
        if stock[1] is not None:
            storeDate = datetimeToDate(stock[1])
            store_bus_days = int(numpy.busday_count(storeDate, endDate)) + 1
            store_all_days = (endDate - storeDate).days + 1
            inDb[str(stock[0])] = { "busdays": store_bus_days, "weekend_days": (store_all_days - store_bus_days)}
        else:
            inDb[str(stock[0])] = days
    print inDb
    return inDb

def datetimeToDate(dateFormat):
    return date(dateFormat.year, dateFormat.month, dateFormat.day)

def updateDB(meanResult, stdevResult):
    for ticker, list in meanResult.iteritems():
        stockMean = list
        stockStdev = stdevResult[ticker]
        success = db.updateMeanAndStdev(ticker, stockMean, stockStdev)
        if not success:
            print ticker + " update failed"
    return True

def runTest():
    startTime = time.time()
    getStockTweets()
    print str(int((time.time() - startTime)*1000)) + " milliseconds"

if __name__ == "__main__":
    runTest()