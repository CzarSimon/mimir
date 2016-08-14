'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { length, color, font } from '../../styles/styles';

export default class Title extends Component {
  render() {
    const { title } = this.props;
    return (
      <View style={styles.container}>
        <Text style={styles.title}>{title}</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: 'stretch',
    marginHorizontal: length.button,
    alignItems: 'center',
    justifyContent: 'center'
  },
  title: {
    color: color.blue,
    fontFamily: font.type.sans.normal,
    fontSize: font.h4
  }
})
