'use strict'
import _ from 'lodash'

export const arrayEquals = (a1, a2) => {
  let i = a1.length
  if (i !== a2.length) return false
  while (i--) {
    if (a1[i] !== a2[i]) return false
  }
  return true
}

export const is_positive = changeStr => (
  (_.startsWith(changeStr, "-")) ? false : true
)

export const format_price_change = change => (
  (is_positive(change) ? "+" : "") + round(change).toString() + "%"
)

export const round = (number, decimals = 2) => parseFloat(number).toFixed(decimals)


export const formatName = (name, forbidden = ['inc', 'corporation', 'plc']) => {
  const words = _.split(name, ' ')
  const lower_words = _.map(words, (word) => _.lowerCase(word))
  for (let word of lower_words) {
    if (forbidden.includes(word)) {
      const formated_name = _.join(_.slice(words, 0, _.findIndex(word, null, 1) + 1), ' ')
      return _.replace(formated_name, /,$/, "")
    }
  }
  return name
}

export const create_clean_title = (title) => {
  const forbidden = ['TickerLens']
  const split_title = _.split(_.trim(title), ' - ')
  const formated_title = (split_title.length < 2) ? split_title : _.join(_.initial(split_title), ' - ')
  const no_double_whitespace = _.replace(formated_title, new RegExp("\\s+", "g"), " ")
  const no_url = _.replace(no_double_whitespace, new RegExp("(https?|ftp):\/\/[\.[a-zA-Z0-9\/\-]+", "g"), "")
  const clean_title =  _.reduce(forbidden, (prev, forb) => _.replace(prev, forb, ""), no_url)
  return clean_title
}

export const create_subject_string = (score_object, max_subjects = 3) => {
  const i_obj = _.invert(score_object)
  const all_subjects = _.map(_.orderBy(_.keys(i_obj), 'desc'), val => i_obj[val])
  const subjects = (all_subjects.length > max_subjects) ? _.take(all_subjects, 3) : all_subjects
  const subject_string = _.join(subjects, ', ')
  return (all_subjects.length <= max_subjects) ? subject_string : subject_string + " ..."
}

export const arr_get_value_by_key = (arr = [], val, key = 'Symbol') => (
  _.find(arr, (obj) => (obj[key] === val))
)

export const formatThousands = num_str => (
  _.replace(num_str, /\B(?=(\d{3})+(?!\d))/g, " ")
)
