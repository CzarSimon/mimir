'use strict';

import React, { Component } from 'react';
import { View, Text, Image, TouchableHighlight, StyleSheet } from 'react-native'
import { length, font, color } from '../../styles/styles';
import Icon from 'react-native-vector-icons/Ionicons';

export default class SearchButton extends Component {
  render() {
    const image = (!this.props.active)
    ? <Icon name='ios-search-outline' size={length.icons.medium} color={color.blue} />
    : <Icon name='ios-close-outline' size={length.icons.large} color={color.yellow} />;

    return (
      <View style={styles.container}>
        <TouchableHighlight onPress={() => this.props.action()}>
          <View style={styles.button}>{image}</View>
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
  button: {
    paddingRight: length.mini
  }
})
