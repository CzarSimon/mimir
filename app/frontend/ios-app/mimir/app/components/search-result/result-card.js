'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, length, font } from '../../styles/styles';

export default class ResultCard extends Component {
  handle_click(ticker) {
    console.log('Adding ticker:', ticker);
    this.props.add_ticker(ticker);
  }

  render() {
    const { name, ticker } = this.props;
    return (
      <View style = {styles.container}>
        <View style = {styles.name_info}>
          <Text>{name}</Text>
          <Text>{ticker}</Text>
        </View>
        <TouchableHighlight
          onPress = { () => this.handle_click(ticker)}>
            <Text style={styles.button_text}>+</Text>
        </TouchableHighlight>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'row',
    alignSelf: 'stretch',
    justifyContent: 'space-between',
    borderColor: color.grey.background,
    padding: length.small,
    marginRight: length.medium,
    borderWidth: 1
  },
  name_info: {
    flex: 3,
    alignSelf: 'stretch'
  },
  button_text: {
    color: color.green,
    fontSize: font.h1,
    fontFamily: font.type.dev,
    paddingHorizontal: length.medium
  }
})
