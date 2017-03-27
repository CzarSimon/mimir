import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'

import UrgencyIndicator from '../watchlist/stock-card/urgency-indicator';
import SearchResultContainer from '../../containers/search-result.container';
import Separator from '../helpers/separator';
import { round, formatName, is_positive, format_price_change } from '../../methods/helper-methods';
import { color, font, length } from '../../styles/styles';

export default class BasicInfo extends Component {
  render() {
    const { Name, PercentChange, LastTradePriceOnly, Currency } = this.props.company;
    const change_style = (is_positive(PercentChange)) ? styles.price_up : styles.price_down;
    const change = format_price_change(PercentChange)
    return (
      <View style={styles.container}>
        <SearchResultContainer />
        <View style={styles.card}>
          <UrgencyIndicator {...this.props.twitter_data} />
          <View>
            <Text style={styles.name}>{formatName(Name)}</Text>
            <View style={styles.price}>
              <Text style={styles.price_text}>{round(LastTradePriceOnly)} {Currency}  </Text>
              <Text style={change_style}>{change}</Text>
            </View>
          </View>
        </View>
        <Separator />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    marginLeft: length.medium,
    marginTop: length.navbar,
    paddingBottom: 0
  },
  card: {
    flexDirection: 'row',
    alignSelf: 'stretch',
    paddingVertical: length.small,
  },
  name: {
    fontFamily: font.type.sans.bold,
    fontSize: font.h4,
    color: color.blue
  },
  price: {
    flexDirection: 'row'
  },
  price_text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  },
  price_up: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    color: color.green
  },
  price_down: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    color: color.red
  }
});
