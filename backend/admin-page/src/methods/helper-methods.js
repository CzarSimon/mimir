import _ from 'lodash'
import { companyTerms, devMode } from '../config'


export const createHeaders = token => {
  return (!devMode) ? new Headers({"Authorizaton": token}) : new Headers()
}


export const createHttpObject = (method, token='', payload=undefined) => {
  const body = JSON.stringify(payload)
  const headers = createHeaders(token)
  return { method, headers, body }
}


export const portraitMode = () => {
  const { innerHeight, innerWidth } = window
  return innerHeight >= innerWidth
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


export const parseWebsite = (err, res) => {
  const { error, result } = handleKGResponse(err, res)
  if (!error) {
    return (result) ? result.url : ""
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
