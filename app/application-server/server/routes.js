const database = require("./database");

module.exports = {
  update_stock_stats: (req, res, conn) => {
    const dict = req.body;
    if (dict) {
      database.insert_stock_data(dict.data, conn);
      res.send('success');
    } else {
      res.send('failure');
    }
  }
}
