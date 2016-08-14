'use strict';
import React, { Component } from 'react';
import { View, StyleSheet, ActivityIndicator } from 'react-native';
import { color } from '../styles/styles';

export default class Loading extends Component {
  render = () => (
    <View style={styles.container}>
      <ActivityIndicator size="large" color={color.blue}/>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center'
  }
})
