'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView } from 'react-native'
import SearchResult from './search-result';
import { length, font, color } from '../../styles/styles';

export default class SearchResults extends Component {
  render() {
    const { results, addTicker } = this.props
    const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2})
    const resultList = ds.cloneWithRows(results);
    return (
      <View style = {styles.container}>
        <ListView
          dataSource = {resultLsist}
          renderHeader = {() => <Text style={styles.header}>Search results</Text>}
          renderRow = {(result) => (
            <SearchResult name={result.name} ticker={result.ticker} addTicker={addTicker} />
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
