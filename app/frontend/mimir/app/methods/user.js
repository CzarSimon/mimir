import randomstring from 'randomstring'
import { persistObject } from './async-storage'

export const newUser = () => ({
  id: randomstring.generate(),
  tickers: ['AAPL', 'SNAP', 'TSLA', 'AMZN', 'MSFT'],
  searchHistory: []
})

export const generateNewUser = () => {
  const user = newUser()
  persistObject('user', user)
  return user
}
