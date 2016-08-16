import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native'
import { color, length, font } from '../../../styles/styles';
import { round } from '../../../methods/helper-methods';
import { startsWith } from 'lodash';


export default class Price extends Component {
  format_change = () => {
    const { change } = this.props;
    if (change) {
      const is_up = startsWith(change, '+');
      return {
        change: ((is_up) ? "+" : "") + round(change).toString() + "%",
        change_style: (is_up) ? styles.price_up : styles.price_down
      }
    } else {
      return {
        change: "-",
        change_style: styles.change_null
      }
    }
  }

  render() {
    const { change_style, change } = this.format_change();
    return (
      <View style={styles.price_info}>
        <Text style={change_style}>{ change }</Text>
        <Text style={styles.price}>{round(this.props.price)} {this.props.currency}</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  price_info: {
    flex: 1,
    alignSelf: 'flex-end',
    alignItems: 'flex-end'
  },
  price_up :{
    fontSize: font.h3,
    color: color.green,
    fontFamily: font.type.sans.normal
  },
  price_down :{
    fontSize: font.h3,
    color: color.red,
    fontFamily: font.type.sans.normal
  },
  change_null: {
    fontSize: font.h3,
    color: color.grey.dark,
    fontFamily: font.type.sans.bold
  },
  price: {
    color: color.grey.dark,
    paddingTop: length.mini,
    fontSize: font.h5,
    fontFamily: font.type.sans.normal
  }
});
