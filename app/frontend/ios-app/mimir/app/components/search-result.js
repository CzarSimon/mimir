'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView } from 'react-native'
import ResultCard from './search-result/result-card';
import { length, font, color } from '../styles/styles';

export default class SearchResult extends Component {
  render() {
    const { results, add_ticker } = this.props,
          ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2}),
          result_list = ds.cloneWithRows(results);
    return (
      <View style = {styles.container}>
        <ListView
          dataSource = {result_list}
          renderHeader = {() => <Text style={styles.header}>Search results</Text>}
          renderRow = {(result) => (
            <ResultCard
              name = {result.name}
              ticker = {result.ticker}
              add_ticker = {add_ticker}
            />
          )}
        />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    alignSelf: 'stretch',
    alignItems: 'stretch',
    marginBottom: length.mini
  },
  header: {
    margin: length.small,
    fontSize: font.h4,
    fontFamily: font.type.sans.normal,
    color: color.blue
  }
})
