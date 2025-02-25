import { AsyncStorage } from 'react-native'

export const USER_ID_KEY = 'mimir/user/id';

export const persist = (key, val) => {
  AsyncStorage.setItem(key, val.toString())
  .catch(err => console.log(err))
}

export const persistObject = (key, obj) => {
  AsyncStorage.setItem(key, JSON.stringify(obj))
  .catch(err => console.log(err))
}

export const retrive = (key, show = false) => {
  return (
    AsyncStorage.getItem(key)
    .then(res => {
      if (show) {
        console.log(res);
      }
      return res;
    })
  );
}

export const retriveObject = (key, show = false) => {
  return (
    AsyncStorage.getItem(key)
    .then(res => {
      if (res !== null) {
        if (show) {
          console.log(JSON.parse(res));
        }
        return JSON.parse(res);
      }
    })
  );
}

export const remove = key => {
  return (
    AsyncStorage.removeItem(key)
    .then(res => console.log("Success:", res))
    .catch(err => console.log("Failure:", err))
  );
}
