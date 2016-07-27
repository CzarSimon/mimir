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
          <Text style={styles.text}>{name}</Text>
          <Text style={styles.text}>{ticker}</Text>
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
  text: {
    fontSize: font.text,
    fontFamily: font.type.sans.normal,
  },
  button_text: {
    color: color.green,
    fontSize: font.h1,
    fontFamily: font.type.sans.normal,
    paddingHorizontal: length.medium
  }
})
