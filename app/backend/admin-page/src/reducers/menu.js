import * as types from '../actions/action-types';

const initalState = {
  selected: 'admin page'
}

const menu = (state = initalState, action = {}) => {
  switch (action.type) {
    case types.SELECT_MENU_ITEM:
      return {
        ...state,
        selected: action.payload.menuItem
      }
    default:
      return state;
  }
}

export default menu;
