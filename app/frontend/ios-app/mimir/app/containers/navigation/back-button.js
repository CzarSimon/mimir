'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { length, color } from '../../styles/styles';
import Icon from 'react-native-vector-icons/Ionicons';

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
        <TouchableHighlight
          onPress={() => this.handleClick(nav)}
          underlayColor={color.grey.background}>
          <View style={styles.container}>
              <View style={styles.button}>
                <Icon name='ios-arrow-back-outline' size={length.icons.medium} color={color.blue} />
              </View>
          </View>
      </TouchableHighlight>
      );
    }
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    alignSelf: 'stretch',
    justifyContent: 'center',
    width: length.button
  },
  button: {
    flex: 1,
    paddingVertical: length.mini + 2 
  }
})
