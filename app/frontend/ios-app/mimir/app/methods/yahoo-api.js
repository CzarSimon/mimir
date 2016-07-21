'use strict'
import { join, zipObject } from 'lodash';
const moment = require('moment');

const simple_data = "Name, Symbol, ChangeinPercent, PercentChange, LastTradePriceOnly, Currency, Volume, Ask, Bid";
const detailed_data = simple_data + ", Open, EBITDA, PERatio, MarketCapitalization, EarningsShare, YearHigh, YearLow, AverageDailyVolume, PreviousClose, ChangeFromYearHigh, ChangeFromYearLow"

export const retrive_stock_data = (tickers) => {
  const data_fields = (tickers.length > 1) ? simple_data : detailed_data;
  const api = "https://query.yahooapis.com/v1/public/yql?q="
  const data = encodeURIComponent("select " + data_fields + " from yahoo.finance.quotes where symbol in ('" + join(tickers, "','") + "')");
  const format = "&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"

  if (tickers.length > 1) {
    return (
      fetch(api + data + format)
      .then(res => res.json())
      .then(json => zipObject(tickers, json.query.results.quote))
    );
  } else {
    return (
      fetch(api + data + format)
      .then(res => res.json())
      .then(json => zipObject(tickers, [json.query.results.quote]))
    );
  }
}

export const retrive_historical_data = (ticker, period = "THREE_MONTHS") => {
  const interval = get_date_interval(period);
  const time_period = "startDate = '" + interval.start + "' and endDate = '" + interval.end + "'";
  const api = "https://query.yahooapis.com/v1/public/yql?q="
  const data = encodeURIComponent("select Adj_Close, Date from yahoo.finance.historicaldata where symbol = '" + ticker + "' and " + time_period);
  const format = "&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"
  return (
    fetch(api + data + format)
    .then(res => res.json())
    .then(json => json.query.results.quote)
  );
}

const get_date_interval = (interval) => {
  const frmt = "YYYY-MM-DD"
  const yesterday = moment().subtract(1, 'day').format(frmt);
  switch (interval) {
    case "THREE_MONTHS":
      return {
        start: moment().subtract(3, 'months').format(frmt),
        end: yesterday
      };
    default:
      return {
        start: moment().subtract(3, 'months').format(frmt),
        end: yesterday
      };
  }
}
