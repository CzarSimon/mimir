import DeviceInfo from 'react-native-device-info'
import { persistObject } from './async-storage'

const generateUserId = () => DeviceInfo.getUniqueID()

export const newUser = () => ({
  id: generateUserId(),
  tickers: ['AAPL', 'SNAP', 'TSLA', 'AMZN', 'MSFT'],
  searchHistory: []
})

export const generateNewUser = () => {
  const user = newUser()
  persistObject('user', user)
  return user
}
