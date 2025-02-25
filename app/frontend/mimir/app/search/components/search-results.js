'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet, ListView } from 'react-native'
import SearchResultContainer from '../containers/search-result'
import { length, font, color } from '../../styles/styles'

export default class SearchResults extends Component {
  render() {
    const { results, goToStock } = this.props
    const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2})
    const resultList = ds.cloneWithRows(results)
    return (
      <View style = {styles.container}>
        <ListView
          dataSource={resultList}
          enableEmptySections={true}
          renderHeader={() => <Text style={styles.header}>Search results</Text>}
          renderRow={resultInfo => (
            <SearchResultContainer resultInfo={resultInfo} goToStock={goToStock} />
          )}
        />
      </View>
    )
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
