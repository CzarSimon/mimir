import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native'
import { color, font, length } from '../../../styles/styles';
import { formatName } from '../../../methods/helper-methods';

export default class Name extends Component {
  render() {
    const { name, ticker } = this.props;
    return (
      <View style = {styles.name_info}>
        <Text style={styles.name}>{ formatName(name) }</Text>
        <Text style={styles.ticker}>{ ticker }</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  name_info: {
    flex: 2,
    alignSelf: 'flex-start',
  },
  name: {
    fontSize: font.h3,
    color: color.black,
    fontFamily: font.type.sans.bold
  },
  ticker: {
    color: color.grey.dark,
    paddingTop: length.mini,
    fontSize: font.text,
    fontFamily: font.type.sans.normal
  }
});
