const database = require("./database");

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

module.exports = {
  updateStockStats,
  getStockData
}
