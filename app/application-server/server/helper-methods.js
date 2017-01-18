'use strict'
const _ = require('lodash');

const parseStockList = (dict) => _.values(dict)

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
