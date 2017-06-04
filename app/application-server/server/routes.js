const database = require("./database");
const postgres = require('./postgres');

const updateStockStats = (req, res, conn) => {
  const dict = req.body;
  if (dict) {
    database.insertStockData(dict.data, conn);
    res.sendStatus(200);
  } else {
    res.sendStatus(500);
  }
}

const getStockData = (req, res, conn) => {
  tickers = req.body.tickers;
  if (tickers.length) {
    database.fetchStockData(tickers, conn, (err, res) => {
      if (err) {
        res.status(500).send({data: null, error: err.message});
      } else {
        res.status(200).send({data: res, error: null})
      }
    })
  }
}

const getSearchSugestions = (req, res, pg) => {
  postgres.getSearchSugestions(parseTickers(req), pg, (err, result) => {
    if (err) {
      console.log(err);
      res.sendStatus(500);
    } else {
      const sugestions = JSON.stringify(result.rows);
      res.status(200).send(sugestions);
    }
  })
}

const parseTickers = request => (
  (request.body.tickers) ? request.body.tickers : request.query.tickers
)

module.exports = {
  updateStockStats,
  getStockData,
  getSearchSugestions
}
