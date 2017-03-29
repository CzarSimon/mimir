'use strict'

import React, { Component } from 'react'
import { View, Text, TextInput, StyleSheet } from 'react-native'
import { trim } from 'lodash'
import { length, color, font } from '../../styles/styles'

export default class SearchBar extends Component {
  handleSubmit = query => {
    const { addQuery, runQuery } = this.props
    addQuery(query)
    runQuery(query)
  }

  runQuery = query => {
    this.props.runQuery(query)
  }

  render() {
    const { query } = this.props
    return (
      <View style={styles.container}>
        <TextInput
          style={styles.searchBox}
          onChangeText={(text) => this.runQuery(text)}
          selectionColor={color.blue}
          clearButtonMode='always'
          returnKeyType='search'
          autoCorrect={false}
          autoFocus={true}
          autoCapitalize='none'
          value={query}
          placeholder={"Search tickers"}
          onSubmitEditing={() => this.handleSubmit(query)}
          onFocus={() => this.runQuery("")}
        />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: 'stretch',
    marginLeft: length.button,
    marginRight: length.medium
  },
  searchBox: {
    flex: 1,
    margin: length.mini + 3,
    paddingLeft: length.medium,
    borderRadius: 3,
    backgroundColor: color.grey.background,
    color: color.blue,
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  }
})
