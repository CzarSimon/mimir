from datetime import datetime
from datetime import date
import json
import requests
import numpy

from database import manager as db
from credentials import server

def rankUrgency(stocks):
    stock_urgency = {}
    date = datetime.utcnow()
    minute = date.minute + 1
    hour = date.hour
    statsAndName = _unpackMeanAndStdev(stocks, hour, date)
    for ticker, volume in stocks.iteritems():
        stock_urgency[ticker] = {
            "name": statsAndName[ticker]["name"],
            "volume": volume,
            "mean": float(statsAndName[ticker]["mean"]),
            "stdev": float(statsAndName[ticker]["stdev"]),
            "minute": minute
        }
    _sendToClients(stock_urgency)
    print "Urgency calculation done"
    return True

def _unpackMeanAndStdev(stocks, hour, curr_date):
    is_weekday = 0 if numpy.is_busday([date(curr_date.year, curr_date.month, curr_date.day)])[0] else 1
    stockList = db.queryDatabase('SELECT ticker, name, mean, stdev FROM stocks',True)
    meanAndStdev = {}
    for stock in stockList:
        mean = stock[2][is_weekday][hour]
        stdev = stock[3][is_weekday][hour]
        meanAndStdev[str(stock[0])] = {"name": str(stock[1]), "mean": mean, "stdev": stdev}
    return meanAndStdev

def _sendToClients(stock_info):
    payload = json.dumps(stock_info)
    url = server["POST_URGENCY_URL"]
    headers = {'content-type': 'application/json'}
    requests.post(url=url, data=payload, headers=headers)
    return True

def printStockVolume(stocks):
    for key, val in stocks.iteritems():
        print str(key) + " mentioned " + str(val) + " this hour"