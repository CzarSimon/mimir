'use strict'
const serverInfo = require('./credentials').server,
      express = require('express'),
      path = require('path'),
      bodyParser = require('body-parser');

const makeStockList = require('./server/helperMethods').makeStockList,
      getDate = require('./server/helperMethods').getDate,
      parseStockList = require('./server/helperMethods').parseStockList;

let   app = express(),
      server = require('http').createServer(app);
const io = require('socket.io').listen(server);

let   stockList = makeStockList(),
      lastUpdated = getDate();

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
    io.sockets.emit('update stocklist', {list: stockList, date: lastUpdated})
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
  socket.on('send info to server', (data) => {
    console.log("New user made connection on: " + data.clientMachine);
  })
});
