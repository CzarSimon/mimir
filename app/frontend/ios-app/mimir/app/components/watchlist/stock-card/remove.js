'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, length, font } from '../../../styles/styles';

export default class Remove extends Component {
  render() {
    const { visable, remove_ticker, ticker } = this.props;
    if (visable) {
      return (
        <View style={styles.container}>
          <TouchableHighlight
            onPress = {() => { remove_ticker(ticker) }}>
            <View style={styles.button}>
              <Text style={styles.button_text}>-</Text>
            </View>
          </TouchableHighlight>
        </View>
      );
    } else {
      return (<View />);
    }
  }
}

const styles = StyleSheet.create({
  container: {
    justifyContent: 'center'
  },
  button: {
    opacity: 0.5,
    borderWidth: 2,
    borderRadius: 50,
    borderColor: color.red,
    width: 35,
    marginHorizontal: length.small,
    alignItems: 'center'
  },
  button_text: {
    fontSize: font.h4,
    fontWeight: 'bold',
    fontFamily: font.type.sans.normal,
    color: color.red,
    padding: length.mini
  }
});
