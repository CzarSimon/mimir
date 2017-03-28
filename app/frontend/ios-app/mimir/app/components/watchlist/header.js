'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { length, color, font } from '../../styles/styles';

export default class Header extends Component {
  render() {
    const { modifiable, handleClick } = this.props,
          buttonText = (!this.props.modifiable) ? "Change" : "Done";
    return (
      <View style = {styles.container}>
        <Text style={styles.headerText}>Watchlist</Text>
        <TouchableHighlight
          onPress = {() => handleClick()}>
          <View style={styles.button}>
            <Text style={styles.buttonText}>{buttonText}</Text>
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
    marginVertical: length.small,
  },
  headerText: {
    fontSize: font.h3,
    fontFamily: font.type.sans.normal,
    color: color.blue
  },
  button: {
    opacity: 0.5,
    borderWidth: 2,
    borderColor: color.blue,
  },
  buttonText: {
    flex: 1,
    alignSelf: 'center',
    justifyContent: 'center',
    fontWeight: 'bold',
    fontSize: font.text,
    fontFamily: font.type.sans.normal,
    color: color.blue,
    padding: length.mini
  }
});
