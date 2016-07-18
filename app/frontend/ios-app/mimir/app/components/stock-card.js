import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'

import UrgencyIndicator from './stock-card/urgency-indicator';
import Name from './stock-card/name';
import Price from './stock-card/price';
import { round, format_name } from '../methods/helper-methods';
import { color, margin, font } from '../styles/styles';
import { company_page_route } from '../routing/routes';

export default class StockCard extends Component {
  handleClick = (ticker, navigator) => {
    navigator.push(company_page_route(ticker));
  }
  render() {
    const { Name: StockName, Symbol, PercentChange, LastTradePriceOnly, twitter_data, nav } = this.props;
    return (
      <TouchableHighlight
        onPress = { () => this.handleClick(Symbol, nav)}>
        <View style={styles.card} >
          <UrgencyIndicator {...twitter_data} />
          <Name name={StockName} ticker={Symbol} />
          <Price change={PercentChange} tic={Symbol} price={LastTradePriceOnly} />
        </View>
      </TouchableHighlight>
    );
  }
}

const styles = StyleSheet.create({
  card: {
    flex: 1,
    flexDirection: 'row',
    alignSelf: 'stretch',
    paddingVertical: margin.small,
    marginHorizontal: margin.medium,
    marginBottom: margin.small,
    borderColor: color.grey.background,
    borderWidth: 1
  }
});
