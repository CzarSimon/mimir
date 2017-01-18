import _ from 'lodash'


export const portraitMode = () => {
  const { innerHeight, innerWidth } = window
  return innerHeight > innerWidth
}


export const objectArrayToObject = (objectArray, key) => {
  const keys = _.map(objectArray, obj => obj[key])
  return _.zipObject(keys, objectArray)
}


export const sortByKey = (arr, key) => _.sortBy(arr, key)
