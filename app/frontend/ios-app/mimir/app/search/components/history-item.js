'use strict'

import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { color, font } from '../../styles/styles'
import Separator from '../../components/helpers/separator'

export default class HistoryItem extends Component {
  render() {
    const { text } = this.props
    return (
      <View>
        <Text style={styles.text}>{text}</Text>
        <Separator customStyles={{backgroundColor: color.grey.medium}}/>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  text: {
    color: color.blue,
    fontSize: font.h5
  }
})
