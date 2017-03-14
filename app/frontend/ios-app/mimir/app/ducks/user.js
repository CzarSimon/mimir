'use strict'
import { createAction } from 'redux-actions'

/* --- Types --- */
export const CLEAR_SEARCH_HISTORY = 'CLEAR_SEARCH_HISTORY'

/* --- Reducer --- */

/* --- Actions --- */
export const clearSearchHistory = createAction(CLEAR_SEARCH_HISTORY)
