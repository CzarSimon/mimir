'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import SafariView from 'react-native-safari-view';
import { color, length, font } from '../../../styles/styles';

export default class ArticleButton extends Component {
  go_to_article = () => {
    SafariView.isAvailable()
    .then(SafariView.show({
      url: this.props.url,
      tintColor: color.blue
    }))
  }

  render() {
    return (
      <TouchableHighlight
        onPress = {() => this.go_to_article()}>
        <View style={styles.button}>
          <Text style={styles.text}>Go to article</Text>
        </View>
      </TouchableHighlight>
    )
  }
}

const styles = StyleSheet.create({
  button: {
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
