'use strict'

const parseStockList = (dict) => {
  let list = [];
  for (let key in dict) {
    if (!dict.hasOwnProperty(key)) continue;
    list.push(dict[key]);
  }
  return list;
}

const nowUTC = () => new Date(new Date().toUTCString().substr(0,25));

const isWeekend = () => {
  const day = nowUTC().getDay();
  return (day === 0) || (day === 6);
}

const dayType = () => (isWeekend()) ? "weekend_days" : "busdays";

module.exports = {
  parseStockList,
  dayType,
  nowUTC,
};
