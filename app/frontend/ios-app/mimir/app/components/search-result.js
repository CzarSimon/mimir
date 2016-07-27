'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView } from 'react-native'
import ResultCard from './search-result/result-card';
import { length } from '../styles/styles';

export default class SearchResult extends Component {
  render() {
    const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2});
    const result_list = ds.cloneWithRows(this.props.results);
    return (
      <View style = {styles.container}>
        <ListView
          dataSource = {result_list}
          renderHeader = {() => <Text style={styles.header}>Search results</Text>}
          renderRow = {(result) => (
            <ResultCard name={result.name} ticker={result.ticker} />
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
    marginBottom: length.medium
  },
  header: {
    marginLeft: length.small,
    marginBottom: length.mini
  }
})
