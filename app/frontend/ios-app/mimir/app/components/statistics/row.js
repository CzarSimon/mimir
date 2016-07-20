'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, length } from '../../styles/styles';

export default class Row extends Component {
  render() {
    const { name, value } = this.props;
    return (
      <View style={styles.container}>
        <Text>{name}</Text>
        <Text style={styles.value}>{value}</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingVertical: length.small,
    paddingRight: length.medium,
    borderBottomWidth: 1,
    borderBottomColor: color.grey.background
  },
  value: {
    fontWeight: 'bold'
  }
})
