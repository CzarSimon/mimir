'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, length, font } from '../../styles/styles'
import Icon from 'react-native-vector-icons/Ionicons'

export default class SearchResult extends Component {
  addTicker = ticker => {
    const { addTicker, goToStock } = this.props
    addTicker(ticker)
    goToStock(ticker)
  }

  render() {
    const { name, ticker } = this.props.resultInfo
    return (
      <View style={styles.container}>
        <View style={styles.nameInfo}>
          <Text style={styles.text}>{name}</Text>
          <Text style={styles.text}>{ticker}</Text>
        </View>
        <TouchableHighlight
          onPress={() => this.addTicker(ticker)}>
          <View style={styles.button}>
            <Icon name='ios-add-circle-outline' size={length.icons.medium} color={color.green} />
          </View>
        </TouchableHighlight>
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
    borderColor: color.grey.background,
    padding: length.small,
    marginRight: length.medium,
    borderWidth: 1,
    backgroundColor: color.white,
    marginBottom: length.small
  },
  nameInfo: {
    flex: 3,
    alignSelf: 'stretch'
  },
  text: {
    fontSize: font.text,
    fontFamily: font.type.sans.normal,
  },
  button: {
    justifyContent: 'center',
    paddingTop: length.mini,
    paddingHorizontal: length.mini
  }
})
