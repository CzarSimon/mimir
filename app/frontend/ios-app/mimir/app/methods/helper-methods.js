'use strict';

const _ = require('lodash');

export const array_equals = (a1, a2) => {
  let i = a1.length;
  if (i !== a2.length) return false;
  while (i--) {
    if (a1[i] !== a2[i]) return false;
  }
  return true;
}

export const is_positive = (change_str) => {
  return (_.startsWith(change_str, "-")) ? false : true
}

export const format_price_change = (change) => {
  return (is_positive(change) ? "+" : "") + round(change).toString() + "%";
}

export const round = (number, decimals = 2) => {
  return parseFloat(number).toFixed(decimals);
}

export const format_name = (name, forbidden = ['inc', 'corporation', 'plc']) => {
  const words = _.split(name, ' ');
  const lower_words = _.map(words, (word) => _.lowerCase(word));

  for (let word of lower_words) {
    if (forbidden.includes(word)) {
      const formated_name = _.join(_.slice(words, 0, _.findIndex(word, null, 1) + 1), ' ');
      return _.replace(formated_name, /,$/, "");
    }
  }
  return name
}

export const create_clean_title = (title) => {
  const forbidden = ['TickerLens'];
  const split_title = _.split(_.trim(title), ' - ');
  const formated_title = (split_title.length < 2) ? split_title : _.join(_.initial(split_title), ' - ');
  const no_double_whitespace = _.replace(formated_title, new RegExp("\\s+", "g"), " ");
  const no_url = _.replace(no_double_whitespace, new RegExp("(https?|ftp):\/\/[\.[a-zA-Z0-9\/\-]+", "g"), "");
  const clean_title =  _.reduce(forbidden, (prev, forb) => _.replace(prev, forb, ""), no_url);
  return clean_title;
}

export const create_subject_string = (score_object) => {
  const i_obj = _.invert(score_object);
  return _.join(_.map(_.orderBy(_.keys(i_obj),'desc'), val => i_obj[val]), ', ');
}

export const arr_get_value_by_key = (arr = [], val, key = 'Symbol') => {
  return _.find(arr, (obj) => (obj[key] === val));
}

export const format_thousands = (num_str) => {
  return _.replace(num_str, /\B(?=(\d{3})+(?!\d))/g, " ");
}
