import * as types from './action-types';
import { createAction } from 'redux-actions';


export const selectMenuItem = createAction(
  types.SELECT_MENU_ITEM, menuItem => ({ menuItem })
)
