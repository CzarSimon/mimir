import _ from 'lodash';
import { companyTerms, kgKey } from '../config'
import { parseCompanyDescription, parseImageUrl, parseWebsite } from './helper-methods';
import { reciveTickerInfo } from '../actions/tickers-actions'
import KGSearch from 'google-kgsearch';


export const getCompanyInfo = (name, ticker, dispatch) => {
  KGSearch(kgKey).search({query: name, limit: 1}, (err, res) => {
    const description = parseCompanyDescription(err, res)
    const imageUrl = parseImageUrl(err, res)
    const website = parseWebsite(err, res)
    dispatch(reciveTickerInfo(ticker, description, name, imageUrl, website))
  })
}

const formatWord = word => _.toLower(word)


export const parseCompanyName = name => {
  const words = _.split(name, ' ');
  for (let i = 0; i < words.length; i++) {
    if (companyTerms.includes(formatWord(words[i]))) {
      return _.join(_.slice(words, 0, i + 1), ' ')
    }
  }
}


export const retriveStockName = ticker => {
  const api = "https://query.yahooapis.com/v1/public/yql?q="
  const data = encodeURIComponent("select Name, Symbol from yahoo.finance.quotes where symbol in ('" + ticker + "')");
  const format = "&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"
  return (
    fetch(api + data + format)
    .then(res => res.json())
    .then(res => res.query.results.quote.Name)
  )
}


export const getNameYahoo = ticker => {
  return retriveStockName(ticker)
  .then(rawName => parseCompanyName(rawName))
}
