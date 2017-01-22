import _ from 'lodash'
import { companyTerms } from '../config'


export const portraitMode = () => {
  const { innerHeight, innerWidth } = window
  return innerHeight > innerWidth
}


export const objectArrayToObject = (objectArray, key) => {
  const keys = _.map(objectArray, obj => obj[key])
  return _.zipObject(keys, objectArray)
}


export const sortByKey = (arr, key) => _.sortBy(arr, key)


export const replaceEscapeChars = text => _.unescape(text)


const handleKGResponse = (err, res) => {
  if (err) console.log(err);
  return (!err && res[0])
    ? {error: false, result: res[0].result}
    : {error: true, result: null}
}


export const parseCompanyDescription = (err, res) => {
  const { error, result } = handleKGResponse(err, res)
  if (!error) {
    return (result.detailedDescription)
      ? replaceEscapeChars(result.detailedDescription.articleBody)
      : ""
  } else {
    return ""
  }
}


export const parseImageUrl = (err, res) => {
  const { error, result } = handleKGResponse(err, res)
  if (!error) {
    return (result.image) ? result.image.contentUrl : ""
  } else {
    return ""
  }
}


const formatWord = word => _.toLower(word)


export const parseCompanyName = name => {
  const words = _.split(name, ' ');
  for (let i = 0; i < words.length; i++) {
    if (companyTerms.includes(formatWord(words[i]))) {
      return _.join(_.slice(words, 0, i + 1), ' ')
    }
  }
  return name
}
