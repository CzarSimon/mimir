import { join, zipObject } from 'lodash';

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
