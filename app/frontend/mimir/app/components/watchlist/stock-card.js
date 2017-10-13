import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import Swipeout from 'react-native-swipeout'
import UrgencyIndicator from './stock-card/urgency-indicator'
import Name from './stock-card/name'
import Price from './stock-card/price'
import Remove from './stock-card/remove'
import { round, formatName } from '../../methods/helper-methods'
import { color, font, length } from '../../styles/styles'
import { card } from '../../styles/common'

export default class StockCard extends Component {
  handleClick = ticker => {
    this.props.navigate(ticker)
  }

  render() {
    const {
      Name: StockName,
      Symbol,
      PercentChange,
      LastTradePriceOnly,
      Currency,
      twitterData,
      modifiable,
      removeTicker
    } = this.props;
    const deleteButton = [
      {
        backgroundColor: color.red,
        color: color.white,
        onPress: () => removeTicker(Symbol),
        component: <Remove />
      }
    ];
    return (
      <View style={styles.card}>
        <Swipeout right={deleteButton} backgroundColor={color.white}>
          <TouchableHighlight
            onPress = { () => this.handleClick(Symbol)}
            underlayColor = {color.grey.background}>
                <View style={styles.cardContents}>
                  <UrgencyIndicator {...twitterData} />
                  <Name name={StockName} ticker={Symbol} />
                  <Price change={PercentChange} tic={Symbol} price={LastTradePriceOnly} currency={Currency} />
                </View>
          </TouchableHighlight>
        </Swipeout>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  cardContents: {
    flex: 1,
    flexDirection: 'row',
    alignSelf: 'stretch',
    paddingVertical: length.medium,
    paddingHorizontal: length.medium,
  },
  card: {
    marginHorizontal: length.medium,
    marginBottom: length.small,
    backgroundColor: color.blue,
    ...card
  }
});
