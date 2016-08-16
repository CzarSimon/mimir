'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, length, font } from '../../../styles/styles';
import Icon from 'react-native-vector-icons/Ionicons';

export default class Remove extends Component {
  render() {
    const { visable, remove_ticker, ticker } = this.props;
    if (visable) {
      return (
        <View style={styles.container}>
          <TouchableHighlight
            onPress = {() => { remove_ticker(ticker) }}
            underlayColor={color.green}>
            <View style={styles.button}>
              <Icon name='ios-remove-circle-outline' size={length.icons.medium} color={color.red} />
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
    justifyContent: 'center',
  },
  button: {
    paddingLeft: length.small,
    alignItems: 'center'
  }
});
