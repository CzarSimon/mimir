'use strict';

import React, { Component } from 'react';
import { View, Text, Image, TouchableHighlight, StyleSheet } from 'react-native'
import { length } from '../../styles/styles';

export default class SearchButton extends Component {
  render() {
    const image = (!this.props.active) ? "Search" : "Close";
    return (
      <View style={styles.container}>
        <TouchableHighlight onPress={() => this.props.action()}>
          <Text>{image}</Text>
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
  }
})
