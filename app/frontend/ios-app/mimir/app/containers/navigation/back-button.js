'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { length } from '../../styles/styles';

export default class BackButton extends Component {
  handleClick = (navigator) => {
    navigator.pop();
  }
  render() {
    const { index, nav } = this.props;
    if (index === 0) {
      return (<View style={styles.container}/>);
    } else {
      return (
        <View style={styles.container}>
          <TouchableHighlight onPress={() => this.handleClick(nav)}>
            <Text>back</Text>
          </TouchableHighlight>
        </View>
      );
    }
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
