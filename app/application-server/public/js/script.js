var printFn = function() {
  alert('I am loaded, ES6 works');
}

var listItem = function(item) {
  var level = urgencyLevel(item.volume, item.mean, item.stdev, item.minute);
  itemId = (item.name).replace(/\s/g, '');
  var hidden_content = "<p>Volume: " + item.volume + "</p><p>Mean: " + item.mean + "</p><p>Stdev: " + item.stdev + "</p>"
  var hidden = "<div class='info' id='" + itemId +"' style='display:none'>" + hidden_content + "</div>"
  var visable_content = "<p class='name'>" + item.name + "</p><p class='urgency'>" + item.urgency + "</p>"
  var visable = "<div onclick=toggleListItem('" + itemId + "')>" + visable_content + "</div>"
  return "<li class='" + level + "'>" + visable + hidden + "</li>";
}

var toggleListItem = function(itemId) {
  var item = document.getElementById(itemId);
  if (item.style.display === 'none') {
    item.style.display = 'block';
  } else {
    item.style.display = 'none';
  }
}

var urgencyLevel = function(volume, mean, stdev, minute) {
  var damping = parseFloat(minute) / 60.0;
  if (volume <= (damping * (mean + stdev))) {
    return "lvl-normal";
  } else if (volume <= (damping * (mean + 2 * stdev))) {
    return "lvl-high";
  } else {
    return "lvl-urgent"
  }
}

var makeList = function(list) {
  var htmlList = "";
  var roundedList = list.map(function(item) {
    var urg;
    if (item.mean > 0) {
      urg = ((60.0 / parseFloat(item.minute)) * parseFloat(item.volume)/parseFloat(item.mean)).toFixed(2);
    } else{
      urg = 0.01;
    }
    return Object.assign({}, item, {
      urgency: urg
    });
  });
  var sortedList = sort(roundedList);
  for (item of sortedList) {
    htmlList = htmlList + listItem(item);
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
