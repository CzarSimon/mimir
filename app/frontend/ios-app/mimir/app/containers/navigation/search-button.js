'use strict';

import React, { Component } from 'react';
import { Text, Image, TouchableHighlight, StyleSheet } from 'react-native'

export default class SearchButton extends Component {
  handleClick = () => {
    console.log('searching');
    //this.props.toggle_search_bar();
  }
  render() {
    const image = (!this.props.active) ? "search" : "close";
    return (
      <TouchableHighlight onPress={() => this.handleClick()}>
        <Text>{image}</Text>
      </TouchableHighlight>
    );
  }
}
