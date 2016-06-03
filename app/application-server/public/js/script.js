var printFn = function() {
  alert('I am loaded, ES6 works');
}

var listItem = function(name, value) {
  var urg = (value < 1.0) ? 'non-urgent' : 'urgent';
  return "<li><p class='name'>" + name + "</p><p class='urgency " + urg + "'>" + value + "</p></li>";
}

var makeList = function(list) {
  var htmlList = "";
  var roundedList = list.map(function(item) {
    return Object.assign({}, item, {
      urgency: parseFloat(item.urgency).toFixed(2)
    });
  });
  var sortedList = sort(roundedList);
  for (item of sortedList) {
    console.log(listItem(item.name, item.urgency));
    htmlList = htmlList + listItem(item.name, item.urgency)
  }
  document.getElementById('stockList').innerHTML = htmlList;
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
  xmlHttp.open("GET", '/stockList', true);
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
