'use strict'
import React, { Component } from 'react'
import { View, Text } from 'react-native'
import ClearHistoryContainer from '../containers/clear-history-button'

export default class HistoryHeader extends Component {
  render() {
    return (
      <View>
        <Text>Search History</Text>
        <ClearHistoryContainer />
      </View>
    )
  }
}
