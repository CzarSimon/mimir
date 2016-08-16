'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, font, length } from '../../styles/styles';

export default class Header extends Component {
  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.header_text}>Top News</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    alignItems: 'flex-start',
    margin: length.medium,
    marginTop: length.small
  },
  header_text: {
    fontSize: font.h3,
    fontFamily: font.type.sans.normal
  }
})
