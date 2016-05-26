import sys
sys.path.append("..")
from database import manager as db
import threading
from datetime import date
import numpy
import math, time, schedule

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
    return True

    # --- THIS WAS REPLACED BY A CALL TO THE DISPATCH METHOD
    # d = threading.Thread(name=str(prevTicker) + " thread", target=reduceByHour, args=(tickerList, prevTicker, days, stockMeans, stockStdevs, ))
    # d.setDaemon(True)
    # threads[prevTicker] = d
    # d.start()

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
    meanList = [0.0]*24
    stdevList = [0.0]*24
    if days > 0:
        for hour, list in hourlyMentions.iteritems():
            mean = sum(list)/float(days)
            stdev = 0
            for dayCount in list:
                stdev += math.pow(float(dayCount) - mean, 2)
            stdev = round(math.sqrt(stdev/(float(days) - 1.0)), 2)
            meanList[hour] = round(mean, 2)
            stdevList[hour] = stdev
    return {"mean": meanList, "stdev": stdevList}


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
    # print ""
    # print "------ Controll -----"
    # controlList = db.queryDatabase('SELECT name, mean, stdev FROM stocks',True)
    # for item in controlList:
    #     print str(item[0]) + " mean is:", item[1]
    #     print str(item[0]) + " stdev is:", item[2]
    #     print "---------------------"

if __name__ == "__main__":
    runTest()
    # schedule.every().day.at('09:22').do(runTest)
    # schedule.every().day.at('21:32').do(runTest)
    # while True:
    #     schedule.run_pending()
    #     pass
