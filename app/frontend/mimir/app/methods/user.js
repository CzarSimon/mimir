import DeviceInfo from 'react-native-device-info'
import SHA256 from 'crypto-js/sha256'
import { persistObject } from './async-storage'

const generateUserId = () => SHA256(DeviceInfo.getUniqueID()).toString()

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
