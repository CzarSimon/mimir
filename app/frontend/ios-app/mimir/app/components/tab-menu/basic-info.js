import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'

import UrgencyIndicator from '../watchlist/stock-card/urgency-indicator';
import SearchResultContainer from '../../containers/search-result.container';
import Separator from '../helpers/separator';
import { round, format_name } from '../../methods/helper-methods';
import { color, margin, font, length } from '../../styles/styles';

export default class BasicInfo extends Component {
  render() {
    const { Name, PercentChange, LastTradePriceOnly, Currency } = this.props.company;
    return (
      <View style={styles.container}>
        <SearchResultContainer />
        <View style={styles.card}>
          <UrgencyIndicator {...this.props.twitter_data} />
          <View style={styles.info}>
            <Text style={styles.name}>{format_name(Name)}</Text>
            <View style={styles.price}>
              <Text>{round(LastTradePriceOnly)} {Currency}</Text>
              <Text>  {PercentChange}</Text>
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
    marginTop: length.navbar
  },
  card: {
    flexDirection: 'row',
    alignSelf: 'stretch',
    paddingBottom: length.small
  },
  info: {
    paddingLeft: length.small
  },
  name: {
    fontWeight: 'bold'
  },
  price: {
    flexDirection: 'row'
  }
});
