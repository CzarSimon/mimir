import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native'
import { color, margin, font } from '../../../styles/styles';
import { format_name } from '../../../methods/helper-methods';

export default class Name extends Component {
  render() {
    const { name, ticker } = this.props;
    return (
      <View style = {styles.name_info}>
        <Text style={styles.name}>{ format_name(name) }</Text>
        <Text style={styles.ticker}>{ ticker }</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  name_info: {
    flex: 2,
    alignSelf: 'flex-start',
    paddingLeft: margin.small,
  },
  name: {
    fontSize: font.h3,
    color: color.black,
    fontFamily: font.type.sans.normal
  },
  ticker: {
    color: color.grey.dark,
    paddingTop: margin.mini,
    fontSize: font.text,
    fontFamily: font.type.sans.normal
  }
});
