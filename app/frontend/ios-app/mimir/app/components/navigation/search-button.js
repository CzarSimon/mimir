'use strict';

import React, { Component } from 'react';
import { View, Text, Image, TouchableHighlight, StyleSheet } from 'react-native'
import { length, font, color } from '../../styles/styles';

export default class SearchButton extends Component {
  render() {
    const image = (!this.props.active)
    ? {text: "Search", style: styles.search_text}
    : {text: "Close", style: styles.close_text};

    return (
      <View style={styles.container}>
        <TouchableHighlight onPress={() => this.props.action()}>
          <Text style={image.style}>{image.text}</Text>
        </TouchableHighlight>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    width: length.button
  },
  search_text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    color: color.blue
  },
  close_text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    color: color.yellow
  }
})
