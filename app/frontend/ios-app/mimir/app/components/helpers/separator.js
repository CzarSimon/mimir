import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native';
import { color, length } from '../../styles/styles';

export default class Separator extends Component {
  render() {
    return <View style={styles.separator}/>
  }
}

const styles = StyleSheet.create({
  separator: {
    alignSelf: 'stretch',
    backgroundColor: color.blue,
    height: 1,
    justifyContent: 'center'
  }
})
