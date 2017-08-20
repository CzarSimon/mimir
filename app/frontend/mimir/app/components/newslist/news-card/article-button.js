'use strict'

import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import SafariView from 'react-native-safari-view';
import { color, length, font } from '../../../styles/styles'
import { card } from '../../../styles/common'

export default class ArticleButton extends Component {
  goToArticle = () => {
    SafariView.isAvailable()
    .then(SafariView.show({
      url: this.props.url,
      tintColor: color.blue
    }))
  }

  render() {
    return (
      <TouchableHighlight
        onPress = {() => this.goToArticle()}>
        <View style={styles.button}>
          <Text style={styles.text}>Go to article</Text>
        </View>
      </TouchableHighlight>
    )
  }
}

const styles = StyleSheet.create({
  button: {
    ...card,
    flex: 1,
    alignSelf: 'stretch',
    backgroundColor: color.blue,
    alignItems: 'center',
    marginBottom: length.mini
  },
  text: {
    fontFamily: font.type.sans.bold,
    fontSize: font.h4,
    paddingVertical: length.small,
    color: color.white
  }
})
