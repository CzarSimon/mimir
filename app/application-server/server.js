'use strict'

const serverInfo = require('./credentials').server,
      express = require('express'),
      path = require('path'),
      bodyParser = require('body-parser'),
      zipObject = require('lodash')['zipObject'],
      map = require('lodash')['map'];

const makeStockList = require('./server/helperMethods').makeStockList,
      getDate = require('./server/helperMethods').getDate,
      parseStockList = require('./server/helperMethods').parseStockList;

const app = express(),
      server = require('http').createServer(app),
      io = require('socket.io').listen(server);

let   stockList = makeStockList(),
      lastUpdated = getDate(),
      stock_dict;

app.use(express.static(path.join(__dirname, 'public')));
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.get('/stockList', (req, res) => {
  const payload = {
    data: stock_dict,
    list: stockList,
    date: lastUpdated
  };
  res.send(payload);
});

app.post('/stockList', (req, res) => {
  const dict = req.body;
  if (dict) {
    stockList = parseStockList(dict);
    lastUpdated = getDate()
    stock_dict = dict;
    io.sockets.emit('update stocklist', {list: stockList, date: lastUpdated})
    io.sockets.emit('NEW TWITTER DATA');
    res.send('success');
  } else {
    res.send('failure');
  }
});

server.listen(serverInfo.port, serverInfo.IP, () => {
  console.log('Server listening on: ' + serverInfo.IP + ":" + serverInfo.port);
});

io.on('connection', (socket) => {
  console.log('connection!');
  socket.emit('get info from client', "GET INFO");
  socket.on('send info to server', data => {
    console.log("New user made connection on: " + data.clientMachine);
  })

  socket.on('GET TWITTER DATA', data => {
    const user_tickers = data.user.tickers;
    if (user_tickers.length) {
      socket.emit('DISPATCH TWITTER DATA', {
        data: zipObject(user_tickers, map(user_tickers, ticker => stock_dict['$' + ticker])),
        date: lastUpdated
      });
    }
  })
  //Have made changes here
  socket.on("FETCH_SEARCH_RESULTS", data => {
    const result = stock_dict['$' + data.query];
    if (result) {
      socket.emit('DISPATCH_SEARCH_RESULTS', {
        results: [Object.assign({}, result, { ticker: data.query })]
      })
    } else {
      socket.emit('DISPATCH_SEARCH_FAILURE', { results: null })
    }
  })
});
