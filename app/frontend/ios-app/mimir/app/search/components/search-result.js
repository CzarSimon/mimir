'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, length, font } from '../../styles/styles';
import Icon from 'react-native-vector-icons/Ionicons';

export default class SearchResult extends Component {
  handleClick = (ticker) => {
    console.log('Adding ticker:', ticker);
    this.props.addTicker(ticker);
  }

  render() {
    const { name, ticker } = this.props;
    return (
      <View style = {styles.container}>
        <View style = {styles.name_info}>
          <Text style={styles.text}>{name}</Text>
          <Text style={styles.text}>{ticker}</Text>
        </View>
        <TouchableHighlight
          onPress = {() => this.handleClick(ticker)}>
            <View style={styles.button}>
              <Icon name='ios-add-circle-outline' size={length.icons.medium} color={color.green} />
            </View>
        </TouchableHighlight>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'row',
    alignSelf: 'stretch',
    justifyContent: 'space-between',
    borderColor: color.grey.background,
    padding: length.small,
    marginRight: length.medium,
    borderWidth: 1,
    backgroundColor: color.white,
    marginBottom: length.mini
  },
  name_info: {
    flex: 3,
    alignSelf: 'stretch'
  },
  text: {
    fontSize: font.text,
    fontFamily: font.type.sans.normal,
  },
  button: {
    justifyContent: 'center',
    paddingTop: length.mini,
    paddingHorizontal: length.mini
  }
})
