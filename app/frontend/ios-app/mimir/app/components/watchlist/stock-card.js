import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'

import UrgencyIndicator from './stock-card/urgency-indicator';
import Name from './stock-card/name';
import Price from './stock-card/price';
import Remove from './stock-card/remove';
import { round, formatName } from '../../methods/helper-methods';
import { color, font, length } from '../../styles/styles';

export default class StockCard extends Component {
  handle_click = (ticker) => {
    this.props.navigate(ticker);
  }

  render() {
    const { Name: StockName, Symbol, PercentChange, LastTradePriceOnly, Currency, twitter_data, modifiable, remove_ticker } = this.props;
    return (
      <TouchableHighlight
        onPress = { () => this.handle_click(Symbol)}
        underlayColor = {color.grey.background}>
        <View style={styles.card}>
          <UrgencyIndicator {...twitter_data} />
          <Name name={StockName} ticker={Symbol} />
          <Price change={PercentChange} tic={Symbol} price={LastTradePriceOnly} currency={Currency} />
          <Remove visable={modifiable} remove_ticker={remove_ticker} ticker={Symbol}/>
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
    paddingVertical: length.small,
    paddingHorizontal: length.medium,
    marginHorizontal: length.medium,
    marginBottom: length.small,
    borderColor: color.grey.background,
    borderWidth: 1,
    borderBottomWidth: 2,
    backgroundColor: color.white
  }
});
