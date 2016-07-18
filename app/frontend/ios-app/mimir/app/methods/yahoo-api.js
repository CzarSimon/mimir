import { join, zipObject } from 'lodash';

export const retrive_stock_data = (tickers) => {
  const api = "https://query.yahooapis.com/v1/public/yql?q="
  const data = encodeURIComponent("select * from yahoo.finance.quotes where symbol in ('" + join(tickers, "','") + "')");
  const format = "&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"

  return (
    fetch(api + data + format)
    .then(res => res.json())
    .then(json => json.query.results.quote)
  );
}
