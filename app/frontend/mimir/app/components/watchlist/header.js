'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { length, color, font } from '../../styles/styles';

export default class Header extends Component {
  render() {
    return (
      <View style = {styles.container}>
        <Text style={styles.headerText}>Watchlist</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginHorizontal: length.medium,
    paddingHorizontal: length.small,
    marginVertical: length.small,
  },
  headerText: {
    fontSize: font.h3,
    fontFamily: font.type.sans.normal,
    color: color.blue
  }
});
