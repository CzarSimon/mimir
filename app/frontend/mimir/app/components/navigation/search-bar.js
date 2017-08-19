'use strict'

import React, { Component } from 'react';
import { View, Text, TextInput, StyleSheet } from 'react-native';
import { length, color, font } from '../../styles/styles';

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
    const { query, keyboardDown, activateSearchKeyboard } = this.props
    const containerStyle = (!keyboardDown) ? styles.containerActive : styles.containerInactive;
    return (
      <View style={containerStyle}>
        <TextInput
          style={styles.searchBox}
          onChangeText={(text) => this.runQuery(text)}
          selectionColor={color.blue}
          clearButtonMode='always'
          returnKeyType='search'
          autoCorrect={false}
          autoFocus={!keyboardDown}
          autoCapitalize='none'
          value={query}
          placeholder={"Search tickers"}
          onFocus={() => activateSearchKeyboard()}
          onSubmitEditing={() => this.handleSubmit(query)}
        />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  containerActive: {
    flex: 1,
    alignSelf: 'stretch',
    marginLeft: length.button,
    marginRight: length.navbar
  },
  containerInactive: {
    flex: 1,
    alignSelf: 'stretch',
    marginLeft: length.button,
    marginRight: length.medium
  },
  searchBox: {
    flex: 1,
    margin: length.mini + 3,
    paddingLeft: length.medium,
    borderRadius: 4,
    backgroundColor: color.grey.background,
    color: color.blue,
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  }
})
