import * as types from '../actions/action-types';

const initalState = {
  username: null,
  token: null
}

const user = (state = initalState, action = {}) => {
  switch (action.type) {
    case types.RECIVE_USER_CREDENTIALS:
      return {
        ...state,
        ...action.payload.user
      };
    case types.LOG_OUT:
      return initalState;
    default:
      return state;
  }
}

export default user;
