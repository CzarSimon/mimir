'use strict';

import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native'
import { color, length, font } from '../../../styles/styles';
import Icon from 'react-native-vector-icons/Ionicons';

export default class Remove extends Component {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.button}>
          <Icon name='ios-remove-circle-outline' size={length.icons.medium} color={color.white} />
        </View>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center'
  },
  button: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center'
  }
});
