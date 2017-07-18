var socket = io.connect();

var TestConnection = function() {
  return "connected";
}

socket.on("GET_CLIENT_INFO", function(data) {
  if (data === 'GET INFO') {
    socket.emit("DISPATCH_CLIENT_INFO", {
      client_machine: navigator.appVersion
    });
  }
})

socket.on("UPDATE_STOCKLIST", function(data) {
  console.log(data);
  makeList(data.list);
  postDate(data.date);
});

const listItem = function(stock) {
  if (!stock.name) return "";
  return (
    "<li class='" + urgencyLevel(stock) + "'>" +
      stockHeader(stock) + stockContent(stock) +
    "</li>"
  )
}

const stockHeader = function(stock) {
  const { name, urgency } = stock;
  const content = "<p class='name'>" + name + "</p><p class='urgency'>" + urgency + "</p>"
  return "<div onclick=toggleListItem('" + createId(name) + "')>" + content + "</div>"
}

const stockContent = function(stock) {
  const { volume, mean, stdev, name } = stock;
  const content = "<p>Volume: " + volume + "</p><p>Mean: " + mean + "</p><p>Stdev: " + stdev + "</p>"
  return "<div class='info' id='" + createId(name) + "' style='display:none'>" + content + "</div>"
}

const createId = function(name) {
  return name.replace(/\s/g, '');
}

var toggleListItem = function(itemId) {
  var item = document.getElementById(itemId);
  if (item.style.display === 'none') {
    item.style.display = 'block';
  } else {
    item.style.display = 'none';
  }
}

const urgencyLevel = function(stock) {
  const { minute, volume, mean, stdev } = stock;
  let damping = parseFloat(minute) / 60.0;
  if (volume <= (damping * (mean + stdev))) {
    return "lvl-normal";
  } else if (volume <= (damping * (mean + 2 * stdev))) {
    return "lvl-high";
  } else {
    return "lvl-urgent"
  }
}

const getStockData = function(stock) {
  dayType = getDayType()
  return Object.assign({}, {
    minute: stock.minute,
    volume: stock.volume,
    mean: stock.mean[dayType][3],
    stdev: stock.stdev[dayType][3]
  })
}

const nowUTC = function() {
  return new Date(new Date().toUTCString().substr(0,25));
}

const getDayType = function() {
  const day = nowUTC().getDay();
  if ((day === 0) || (day === 6)) {
    return "weekend_days"
  } else {
    return "busdays"
  }
}

const calcUrgency = function(stock) {
  const { minute, volume, mean } = getStockData(stock);
  if (mean > 0) {
    const increase = 60.0 / parseFloat(minute);
    const fraction = parseFloat(volume)/parseFloat(mean);
    return (increase * fraction).toFixed(2);
  } else {
    return 0.02
  }
}

const parseStockList = function(stockList) {
  return stockList.map(function(stock) {
    return formatStockObject(stock);
  })
}

const formatStockObject = function(stock) {
  const { minute, volume, mean, stdev } = getStockData(stock);
  return {
    name: stock.name,
    urgency: calcUrgency(stock),
    volume,
    minute,
    mean,
    stdev,
  }
}

var makeList = function(stockList) {
  const sortedStocks = sort(parseStockList(stockList))
  const stocksHtml = sort(sortedStocks).map(function(stock) {
    return listItem(stock)
  }).join("")
  document.getElementById('stockList').innerHTML = stocksHtml;
}

var makeRequest = function() {
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange = function() {
    if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
      var response = JSON.parse(xmlHttp.responseText);
      makeList(response.list);
      postDate(response.date);
    }
  }
  xmlHttp.open("GET", '/stocks', true);
  xmlHttp.send(null);
}

var postDate = function(newDate) {
  var datePlace = document.getElementById('datePlace');
  datePlace.innerHTML = newDate;
}

var sort = function(list) {
  return sortedList = list.sort(function(a,b) {
    var x = a.urgency;
    var y = b.urgency;
    return ((x > y) ? -1 : ((x < y) ? 1 : 0));
  });
}

window.onload = function() {
  makeRequest();
}
