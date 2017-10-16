import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, font, length } from '../../styles/styles';

export default class NoNews extends Component {
  render() {
    return (
      <View style={style.container}>
        <Text style={style.headerText}>There are no current news</Text>
      </View>
    );
  }
}

const style = StyleSheet.create({
  container: {
    alignItems: 'flex-start',
    margin: length.medium,
    marginTop: length.small
  },
  headerText: {
    fontSize: font.h4,
    fontFamily: font.type.sans.normal,
    marginLeft: length.small
  }
})
