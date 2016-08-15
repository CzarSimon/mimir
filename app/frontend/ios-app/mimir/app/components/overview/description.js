'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { round, format_name } from '../../methods/helper-methods';
import { length, font, color } from '../../styles/styles';

export default class Description extends Component {
  render() {
    const { description } = this.props;
    return (
      <View style={styles.container}>
        {
          (description) ? (
            <Text style={styles.text}>{description}</Text>
          ) : (
            <TouchableHighlight
              onPress={() => {console.log("Adding a description")}}>
              <Text style={styles.button_text}>Add a description</Text>
            </TouchableHighlight>
          )
        }
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    marginVertical: length.small,
    marginRight: length.medium
  },
  text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  },
  button_text: {
    alignSelf: 'center',
    fontFamily: font.type.sans.bold,
    fontSize: font.text,
    color: color.yellow,
    borderColor: color.yellow,
    borderWidth: 2,
    padding: length.mini
  }
})
