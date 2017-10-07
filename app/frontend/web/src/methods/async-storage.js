import store from 'store';

//"ee617407-85d4-43bc-9f13-321efdae75e1"
export const USER_ID_KEY = 'mimir/user/id';

export const persist = (key, value) => {
  store.set(key, value);
}

export const persistObject = (key, object) => {
  store.set(key, object);
}

export const retrive = (key, show = false) => {
  const value = store.get(key);
  if (show) {
    console.log(value);
  }
  return value;
}

export const retriveObject = (key, show = false) => {
  const object = store.get(key);
  if (show) {
    console.log(object);
  }
  return object;
}

export const remove = key => {
  store.remove(key);
}
