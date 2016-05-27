from datetime import datetime
from datetime import date
import json
import requests
import numpy

from database import manager as db

def rankUrgency(stocks):
    stockUrgency = {}
    date = datetime.utcnow()
    minute = date.minute
    hour = date.hour
    statsAndName = _unpackMeanAndStdev(stocks, hour, date)
    for ticker, volume in stocks.iteritems():
        threshold = _getThreshold(minute, statsAndName[ticker]["mean"], statsAndName[ticker]["stdev"], 1)
        if float(volume) > threshold:
            print "!!!!!!!! " + ticker + "'s volume above threshold"
        if threshold > 0:
            urgency = float(volume)/threshold
        else:
            urgency = 0.5
        stockUrgency[ticker] = {"name": statsAndName[ticker]["name"], "urgency": urgency}
    _sendToClients(stockUrgency)
    print "Urgency calculation done"
    return True

def _getThreshold(minute, mean, stdev, k):
    return (minute/60.0) * float(mean + k * stdev)

def _unpackMeanAndStdev(stocks, hour, curr_date):
    # {"tickerName":{"name": "stockName", "mean": meanFloat, "stdev": stdevFloat}}
    is_weekday = 0 if numpy.is_busday([date(curr_date.year, curr_date.month, curr_date.day)])[0] else 1
    stockList = db.queryDatabase('SELECT ticker, name, mean, stdev FROM stocks',True)
    meanAndStdev = {}
    for stock in stockList:
        mean = stock[2][is_weekday][hour]
        stdev = stock[3][is_weekday][hour]
        meanAndStdev[str(stock[0])] = {"name": str(stock[1]), "mean": mean, "stdev": stdev}
    return meanAndStdev

def _sendToClients(stockUrgency):
    payload = json.dumps(stockUrgency)
    url = "http://simonlindgren.tech/stockList"
    headers = {'content-type': 'application/json'}
    requests.post(url=url, data=payload, headers=headers)
    return True

def printStockVolume(stocks):
    for key, val in stocks.iteritems():
        print str(key) + " mentioned " + str(val) + " this hour"
