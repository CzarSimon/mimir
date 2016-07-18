'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet,TouchableHighlight } from 'react-native'

export default class BackButton extends Component {
  handleClick = (navigator) => {
    navigator.pop();
  }
  render() {
    const { index, nav } = this.props;
    if (index === 0) {
      return (<View />);
    } else {
      return (
        <TouchableHighlight onPress={() => this.handleClick(nav)}>
          <Text>back</Text>
        </TouchableHighlight>
      );
    }
  }
}
