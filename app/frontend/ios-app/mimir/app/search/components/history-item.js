'use strict'

import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, font, length } from '../../styles/styles'
import Separator from '../../components/helpers/separator'

export default class HistoryItem extends Component {
  render() {
    const { text, handleClick } = this.props
    return (
      <View style={styles.container}>
        <TouchableHighlight onPress={() => handleClick()}>
          <View>
            <Text style={styles.text}>{text}</Text>
            <Separator customStyles={{backgroundColor: color.grey.medium}}/>
          </View>
        </TouchableHighlight>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    marginLeft: length.small
  },
  text: {
    color: color.blue,
    fontSize: font.h4,
    paddingVertical: length.medium
  }
})
