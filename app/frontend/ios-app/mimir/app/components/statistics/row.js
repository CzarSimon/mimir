'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, length, font } from '../../styles/styles';

export default class Row extends Component {
  render() {
    const { name, value } = this.props;
    const component = (value !== 'NaN') ? (
      <View style={styles.container}>
        <Text style={styles.name}>{name}</Text>
        <Text style={styles.value}>{value}</Text>
      </View>
    ) : <View />
    return component;
  }
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: length.small,
    marginHorizontal: length.medium,
    marginBottom: length.mini,
    borderWidth: 1,
    backgroundColor: color.white,
    borderColor: color.grey.background
  },
  value: {
    fontFamily: font.type.sans.bold,
    fontSize: font.text
  },
  name: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  }
})
