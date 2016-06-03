'use strict'
const serverInfo = require('./credentials').server;
const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');
const makeStockList = require('./server/helperMethods').makeStockList;
const getDate = require('./server/helperMethods').getDate;
const parseStockList = require('./server/helperMethods').parseStockList;

let app = express();
let stockList = makeStockList();
let lastUpdated = getDate();

app.use(express.static(path.join(__dirname, 'public')));
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.get('/stockList', (req, res) => {
  let payload = {list: stockList, date: lastUpdated};
  res.send(payload);
});

app.post('/stockList', (req, res) => {
  let dict = req.body;
  if (dict) {
    stockList = parseStockList(dict);
    lastUpdated = getDate();
    res.send('success');
  } else {
    res.send('failure');
  }
});

app.listen(serverInfo.port, serverInfo.IP, () => {
  console.log('Server listening on: ' + serverInfo.IP + ":" + serverInfo.port);
});
