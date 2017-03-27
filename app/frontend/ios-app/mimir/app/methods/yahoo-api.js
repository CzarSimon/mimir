'use strict'
import { join, zipObject } from 'lodash';
const moment = require('moment');

const simpleData = "Name, Symbol, ChangeinPercent, PercentChange, LastTradePriceOnly, Currency, Volume, Ask, Bid";
const detailedData = simpleData + ", Open, EBITDA, PERatio, MarketCapitalization, EarningsShare, YearHigh, YearLow, AverageDailyVolume, PreviousClose, ChangeFromYearHigh, ChangeFromYearLow"

export const retriveStockData = tickers => {
  console.log("yhoo api", tickers)
  const dataFields = (tickers.length > 1) ? simpleData : detailedData;
  const api = "https://query.yahooapis.com/v1/public/yql?q="
  const data = encodeURIComponent("select " + dataFields + " from yahoo.finance.quotes where symbol in ('" + join(tickers, "','") + "')");
  const format = "&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"

  if (tickers.length > 1) {
    return (
      fetch(api + data + format)
      .then(res => res.json())
      .then(json => zipObject(tickers, json.query.results.quote))
      .then(res => {
        console.log("yhoo api",res)
        return res
      })
    );
  } else {
    return (
      fetch(api + data + format)
      .then(res => res.json())
      .then(json => zipObject(tickers, [json.query.results.quote]))
    );
  }
}

export const retriveHistoricalData = (ticker, period = "THREE_MONTHS") => {
  const interval = getDateInterval(period);
  const timePeriod = "startDate = '" + interval.start + "' and endDate = '" + interval.end + "'";
  const api = "https://query.yahooapis.com/v1/public/yql?q="
  const data = encodeURIComponent("select Adj_Close, Date from yahoo.finance.historicaldata where symbol = '" + ticker + "' and " + timePeriod);
  const format = "&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"
  return (
    fetch(api + data + format)
    .then(res => res.json())
    .then(json => json.query.results.quote)
  );
}

const getDateInterval = (interval) => {
  const frmt = "YYYY-MM-DD"
  const yesterday = moment().subtract(1, 'day').format(frmt)
  switch (interval) {
    case "THREE_MONTHS":
      return {
        start: moment().subtract(3, 'months').format(frmt),
        end: yesterday
      };
    case "ONE_WEEK":
      return {
        start: moment().subtract(7, 'days').format(frmt),
        end: yesterday
      }
    default:
      return {
        start: moment().subtract(3, 'months').format(frmt),
        end: yesterday
      };
  }
}
