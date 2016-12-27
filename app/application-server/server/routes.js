const database = require("./database");

const update_stock_stats = (req, res, conn) => {
  const dict = req.body;
  if (dict) {
    database.insert_stock_data(dict.data, conn);
    res.send('success');
  } else {
    res.send('failure');
  }
}

const get_stock_data = (req, res, conn) => {
  tickers = req.body.tickers;
  if (tickers.length) {
    database.fetch_stock_data(tickers, conn, (err, res) => {
      if (err) {
        res.send({data: null, error: err.message});
      } else {
        res.send({data: res, error: null})
      }
    })
  }
}

module.exports = {
  update_stock_stats,
  get_stock_data
}
