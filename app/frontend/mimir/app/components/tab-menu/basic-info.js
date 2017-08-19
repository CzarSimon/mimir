import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'

import UrgencyIndicator from '../watchlist/stock-card/urgency-indicator'
import Separator from '../helpers/separator'
import { round, formatName, is_positive, format_price_change } from '../../methods/helper-methods'
import { color, font, length } from '../../styles/styles'

export default class BasicInfo extends Component {
  render() {
    const { Name, PercentChange, LastTradePriceOnly, Currency } = this.props.company
    const changeStyle = (is_positive(PercentChange)) ? styles.priceUp : styles.priceDown;
    const change = format_price_change(PercentChange)
    return (
      <View style={styles.container}>
        <View style={styles.card}>
          <UrgencyIndicator {...this.props.twitterData} />
          <View>
            <Text style={styles.name}>{formatName(Name)}</Text>
            <View style={styles.price}>
              <Text style={styles.priceText}>{round(LastTradePriceOnly)} {Currency}  </Text>
              <Text style={changeStyle}>{change}</Text>
            </View>
          </View>
        </View>
        <Separator />
      </View>
    )
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
  priceText: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  },
  priceUp: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    color: color.green
  },
  priceDown: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    color: color.red
  }
})
