import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native'
import { color, margin, font } from '../../styles/styles';
import { round } from '../../methods/helper-methods';
import { startsWith } from 'lodash';


export default class Price extends Component {
  format_change = () => {
    const is_up = startsWith(this.props.change, '+');
    return {
      change: ((is_up) ? "+" : "") + round(this.props.change).toString() + "%",
      change_style: (is_up) ? styles.price_up : styles.price_down
    }
  }

  render() {
    const { change_style, change } = this.format_change();
    return (
      <View style={styles.price_info}>
        <Text style={change_style}>{ change }</Text>
        <Text style={styles.price}>{ round(this.props.price) }</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  price_info: {
    flex: 1,
    alignSelf: 'flex-end',
    alignItems: 'flex-end',
    paddingRight: margin.small,
  },
  price_up :{
    fontSize: font.h3,
    color: color.green
  },
  price_down :{
    fontSize: font.h3,
    color: color.red
  },
  price: {
    color: color.grey.dark,
    paddingTop: margin.mini,
    fontSize: font.h5,
    fontFamily: font.dev
  }
});
