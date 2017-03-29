import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { length, font, color } from '../../styles/styles'

export default class ClearHistoryButton extends Component {
  render() {
    return (
      <View>
        <TouchableHighlight
          onPress={() => this.props.clearHistory()}>
          <Text style={styles.text}>Clear</Text>
        </TouchableHighlight>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  text: {
    fontSize: font.h5,
    fontFamily: font.type.sans.normal,
    color: color.grey.dark,
    marginRight: length.mini
  }
})
