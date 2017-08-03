'use strict'
const _ = require('lodash');

const parseStockList = dict => _.values(dict)

// nowUTC() Returns current timestamp a in timezone UTC as a date object
const nowUTC = () => new Date(nowUTCString());

// nowUTCString() Returns current timestamp in timezone UTC as a string
const nowUTCString = () => new Date().toUTCString().substr(0,25);

// isWeekend() Determines if day is a weekend or not
const isWeekend = () => {
  const day = nowUTC().getDay();
  return (day === 0) || (day === 6);
}

// dayType() Returns "weekend_days" if current day is a weekend, "busdays" if not
const dayType = () => (isWeekend()) ? "weekend_days" : "busdays";

// isEmpty() Checks if the supplied argument is null or underfined
const isEmpty = param => (param == null || param == undefined)

module.exports = {
  parseStockList,
  dayType,
  nowUTC,
  nowUTCString,
  isEmpty
};
