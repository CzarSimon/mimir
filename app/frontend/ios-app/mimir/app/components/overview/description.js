'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { round, format_name } from '../../methods/helper-methods';

export default class Description extends Component {
  render() {
    const { description } = this.props;
    if (description) {
      return (
        <View>
          <Text>{description}</Text>
        </View>
      );
    } else {
      return (
        <TouchableHighlight
          onPress={() => {console.log("Adding a description")}}>
          <Text>Add a description</Text>
        </TouchableHighlight>
      )

    }
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center'
  }
})
