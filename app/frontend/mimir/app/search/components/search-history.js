'use strict'
import React, { Component } from 'react'
import { View, Text, FlatList } from 'react-native'
import HistoryHeader from './history-header'
import HistoryItemContainer from '../containers/history-item'

export default class SearchHistory extends Component {
  render() {
    const { history } = this.props;
    return (
      <View>
        <FlatList
          data={history}
          keyExtractor={(item, index) => index}
          ListHeaderComponent={() => <HistoryHeader />}
          renderItem={({ item }) => <HistoryItemContainer text={item} />}
        />
      </View>
    );
  }
}
