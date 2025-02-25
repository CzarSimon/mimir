import _ from 'lodash';
import moment from 'moment';
import sha256 from 'node-forge/lib/sha256';

export const arrayEquals = (a1, a2) => {
  let i = a1.length
  if (i !== a2.length) return false
  while (i--) {
    if (a1[i] !== a2[i]) return false
  }
  return true
}

export const isPositive = number => (number < 0) ? false : true

export const formatPriceChange = change => (
  (isPositive(change) ? "+" : "") + round(change * 100).toString() + "%"
)

export const portraitMode = () => {
  const { innerHeight, innerWidth } = window
  return innerHeight >= innerWidth
}

export const round = (number, decimals = 2) => parseFloat(number).toFixed(decimals)

// isEmpty() Checks if candidate is undefined or null, returns true if so
export const isEmpty = candidate => (
  (candidate === undefined) || (candidate === null)
);

export const companyEndings = ['inc', 'corporation', 'plc', 'company']

export const formatName = (name, forbidden = companyEndings) => {
  const words = _.split(name, ' ')
  const lowerWords = _.map(words, (word) => _.lowerCase(word))
  for (let word of lowerWords) {
    if (forbidden.includes(word)) {
      const formatedName = _.join(_.slice(words, 0, _.findIndex(word, null, 1) + 1), ' ')
      return _.replace(formatedName, /,$/, "")
    }
  }
  return name
}

export const removeStockType = name => {
  const stockType = "Common Stock"
  const transform = (stockName, forbidden) => _.replace(stockName, forbidden, "")
  const functions = [_.capitalize, _.toLower, _.toUpper]
  const cleanName = _.reduce(functions, (name, func) => (
    transform(name, func(stockType))
  ), transform(name, stockType))
  return _.trim(cleanName)
}


export const createCleanTitle = (title) => {
  const forbidden = ['TickerLens']
  const splitTitle = _.split(_.trim(title), ' - ')
  const formatedTitle = (splitTitle.length < 2) ? splitTitle : _.join(_.initial(splitTitle), ' - ')
  const noDoubleWhitespace = _.replace(formatedTitle, new RegExp("\\s+", "g"), " ")
  const noUrl = _.replace(noDoubleWhitespace, new RegExp("(https?|ftp):\/\/[\.[a-zA-Z0-9\/\-]+", "g"), "")
  return _.reduce(forbidden, (prev, forb) => _.replace(prev, forb, ""), noUrl)
}

export const createSubjectString = (score_object, max_subjects = 3) => {
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

export const formatDate = timestamp => moment(timestamp).format('YYYY-MM-DD')

export const hash = message => {
  let messageDigest = sha256.create();
  messageDigest.update(message);
  return messageDigest.digest().toHex();
}
