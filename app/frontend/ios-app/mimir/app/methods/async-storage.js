import { AsyncStorage } from 'react-native';

export const persist = (key, val) => {
  AsyncStorage.setItem(key, val.toString())
  .catch(err => console.log(err))
  .done(); //remove
}

export const persist_object = (key, obj) => {
  console.log(JSON.stringify(obj));
  AsyncStorage.setItem(key, JSON.stringify(obj))
  .catch(err => console.log(err))
  .done(); //remove
}

export const retrive = (key, show = false) => {
  return (
    AsyncStorage.getItem(key)
    .then(res => {
      if (res !== null) {
        if (show) {
          console.log(res);
        }
        return res;
      }
    })
  );
}

export const retrive_object = (key, show = false) => {
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
