'use strict'

import React, { Component } from 'react';
import { View, Text, Image, TouchableHighlight, StyleSheet } from 'react-native';
import { length, font, color } from '../../styles/styles';
import Icon from 'react-native-vector-icons/Ionicons';
import SearchCancel from './search-cancel';

export default class SearchButton extends Component {
  render() {
    const { active, goToSearch, cancelSearch, keyboardDown } = this.props;
    if (!active) {
      return (
        <View style={styles.container}>
          <TouchableHighlight onPress={() => goToSearch()}>
            <View style={styles.button}>
              <Icon name='ios-search-outline' size={length.icons.medium} color={color.blue} />
            </View>
          </TouchableHighlight>
        </View>
      );
    } else if (active && !keyboardDown) {
      return <SearchCancel cancelSearch={cancelSearch}/>
    } else {
      return <View />
    }
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    width: length.button
  },
  button: {
    paddingRight: length.mini
  }
})
