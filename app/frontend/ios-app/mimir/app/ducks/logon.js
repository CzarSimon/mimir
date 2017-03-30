'use strict'
import { retriveObject, persistObject } from '../methods/async-storage'
import { generateNewUser } from '../methods/user'
import { reciveUser, createNewUser } from './user'
import { fetchTwitterData } from './twitter-data'
import { fetchStockData } from './stocks'

/* --- Types --- */
export const USER_LOGON = 'USER_LOGON'

/* --- Reducer --- */

/* --- Actions --- */
export const logonUser = socket => (
  dispatch => (
    retriveObject("user")
    .then(user => {
      const tickers = user.tickers;
      dispatch(reciveUser(user))
      dispatch(fetchTwitterData(user, socket))
      dispatch(fetchStockData(tickers))
    })
    .catch(err => {
      const newUser = generateNewUser()
      persistObject("user", newUser)
      dispatch(createNewUser(newUser))
      dispatch(fetchTwitterData(newUser, socket))
      dispatch(fetchStockData(newUser.tickers))
    })
  )
)
