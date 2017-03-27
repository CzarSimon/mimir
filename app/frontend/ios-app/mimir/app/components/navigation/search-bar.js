'use strict';

import React, { Component } from 'react';
import { View, Text, TextInput, StyleSheet } from 'react-native';
import { trim } from 'lodash';
import { length, color, font } from '../../styles/styles';

export default class SearchBar extends Component {
  constructor(props) {
    super(props);
    this.state = {
      text: null
    };
  }

  handleSubmit = (query = this.state.text) => {
    this.props.runQuery(query);
  }

  handleNewText = (newText) => {
    if (newText.length > 0) {
      this.props.runQuery(newText)
    }
    this.setState({
      text: newText
    })
  }

  render() {
    const placeholder = "Search tickers"
    return (
      <View style={styles.container}>
        <TextInput
          style={styles.search_box}
          onChangeText={(text) => this.handleNewText(text)}
          selectionColor={color.blue}
          clearButtonMode='always'
          returnKeyType='search'
          autoCorrect={false}
          autoFocus={true}
          autoCapitalize='none'
          placeholder={placeholder}
          onSubmitEditing={() => this.handleSubmit()}
          onFocus={() => this.setState({text: ""})}
        />
      </View>
      );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: 'stretch',
    marginHorizontal: length.button
  },
  search_box: {
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
