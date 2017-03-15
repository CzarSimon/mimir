'use strict'
import React, { Component } from 'react'
import { View, Text, ListView } from 'react-native'
import HistoryHeader from './history-header'
import HistoryItemContainer from '../containers/history-item'

export default class SearchHistory extends Component {
  render() {
    const { history } = this.props
    const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2})
    const historyList = ds.cloneWithRows(history)
    return (
      <View>
        <ListView
          dataSource={historyList}
          renderHeader={() => <HistoryHeader />}
          renderRow={itemText => <HistoryItemContainer text={itemText} />}
        />
      </View>
    )
  }
}
