'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { round, formatName } from '../../methods/helper-methods'
import { length, font, color } from '../../styles/styles'
import { card } from '../../styles/common'

export default class Description extends Component {
  render() {
    const { description } = this.props;
    if (description) {
      return (
        <View style={styles.container}>
          <Text style={styles.text}>{description}</Text>
        </View>
      );
    } else {
      return (<View />);
    }
  }
}

const styles = StyleSheet.create({
  container: {
    marginHorizontal: length.medium,
    marginVertical: length.small,
    borderRadius: 5,
    ...card
  },
  text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    padding: length.medium
  },
  button: {
    padding: length.small
  },
  buttonText: {
    alignSelf: 'center',
    fontFamily: font.type.sans.bold,
    fontSize: font.text,
    color: color.yellow,
    borderColor: color.yellow,
    borderWidth: 2,
    padding: length.mini
  }
})
