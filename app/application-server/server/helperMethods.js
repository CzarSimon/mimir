'use strict'

const randomUrgency = () => {
  return (Math.random() * 1.2).toFixed(2);
}

const makeStockList = () => {
  let newList = [
    {name: 'Twitter Inc.', minute: 60, volume: 26, mean: 25, stdev: 5  },
    {name: 'Amazon.com', minute: 60, volume: 24, mean: 25, stdev: 5  },
    {name: 'Apple', minute: 60, volume: 5, mean: 25, stdev: 5  },
    {name: 'LinkedIn', minute: 60, volume: 31, mean: 25, stdev: 5  },
    {name: 'Facebook', minute: 60, volume: 37, mean: 25, stdev: 5  },
    {name: 'Wal-Mart Stores Inc.', minute: 60, volume: 0, mean: 0, stdev: 0  },
    {name: 'The Goldman Sachs Group, Inc.', minute: 60, volume: 37, mean: 0, stdev: 5  },
  ];
  return newList;
}

const parseStockList = (dict) => {
  let list = [];
  for (let key in dict) {
    if (!dict.hasOwnProperty(key)) continue;
    list.push(dict[key]);
  }
  return list;
}

const getDate = () => {
  let date = new Date();
  return date.toLocaleString();
}

module.exports = {
  makeStockList: makeStockList,
  getDate: getDate,
  randomUrgency: randomUrgency,
  parseStockList: parseStockList
};
