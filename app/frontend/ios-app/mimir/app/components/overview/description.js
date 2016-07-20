'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { round, format_name } from '../../methods/helper-methods';
import { length } from '../../styles/styles';

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
              <Text>Add a description</Text>
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
    textAlign: 'justify'
  }
})
