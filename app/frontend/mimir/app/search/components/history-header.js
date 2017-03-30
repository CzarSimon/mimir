'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { length, font, color } from '../../styles/styles'
import ClearHistoryContainer from '../containers/clear-history-button'

export default class HistoryHeader extends Component {
  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.header}>Search History</Text>
        <ClearHistoryContainer />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'row',
    alignSelf: 'stretch',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginTop: length.small,
    marginRight: length.medium
  },
  header: {
    fontSize: font.h4,
    fontFamily: font.type.sans.normal,
    color: color.grey.dark,
    marginLeft: length.small
  }
})
