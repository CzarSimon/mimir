'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { length, color, font } from '../../styles/styles';

export default class Header extends Component {
  render() {
    const { modifiable, handle_click } = this.props,
          button_text = (this.props.modifiable) ? "Done" : "Change";
    return (
      <View style = {styles.container}>
        <Text style={styles.header_text}>Watchlist</Text>
        <TouchableHighlight
          onPress = {() => handle_click()}>
          <View style={styles.button}>
            <Text style={styles.button_text}>{button_text}</Text>
          </View>
        </TouchableHighlight>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginHorizontal: length.medium,
    paddingHorizontal: length.small,
    marginBottom: length.small,
  },
  header_text: {
    fontSize: font.h3,
    fontFamily: font.type.dev,
    color: color.grey.dark
  },
  button: {
    opacity: 0.5,
    borderWidth: 2,
    borderColor: color.grey.dark,
  },
  button_text: {
    flex: 1,
    alignSelf: 'center',
    justifyContent: 'center',
    fontWeight: 'bold',
    fontSize: font.h5,
    fontFamily: font.type.dev,
    color: color.grey.dark,
    padding: length.mini
  }
});
