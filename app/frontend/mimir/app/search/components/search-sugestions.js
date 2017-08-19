'use strict'

import React, { Component } from 'react';
import { length, color, font } from '../../styles/styles';
import { View, Text, FlatList, StyleSheet } from 'react-native';
import SearchSugestion from './search-sugestion';

export default class SearchSugestions extends Component {
  renderHeader = () => (
    <Text style={styles.header}>Sugestions</Text>
  )

  render() {
    const { sugestions, updateAndRunQuery } = this.props
    return (
      <View>
        <FlatList
          data={sugestions}
          renderItem={({ item }) => (
            <SearchSugestion
              sugestion={item}
              updateAndRunQuery={updateAndRunQuery}
            />
          )}
          keyExtractor={item => item.ticker}
          ListHeaderComponent={this.renderHeader}
        />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  header: {
    color: color.blue,
    fontSize: font.h3,
    fontFamily: font.type.sans.bold,
    alignSelf: 'center',
    marginVertical: length.medium
  }
});
