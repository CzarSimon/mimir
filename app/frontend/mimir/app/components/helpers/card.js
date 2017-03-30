import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native';
import { color, length } from '../../styles/styles.js';

export default class Card extends Component {
  render() {
    //Should add children in some way
    return (
      <View style={styles.card}>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  card: {
    marginBottom: length.small,
    marginHorizontal: length.medium,
    backgroundColor: color.white,
    borderWidth: 1,
    borderColor: color.grey.background,
    padding: length.small
  }
})
