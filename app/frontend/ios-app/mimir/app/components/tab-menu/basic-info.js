import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'

import UrgencyIndicator from '../stock-card/urgency-indicator';
import { round, format_name } from '../../methods/helper-methods';
import { color, margin, font } from '../../styles/styles';

export default class BasicInfo extends Component {
  render() {
    const { Name, PercentChange, LastTradePriceOnly, Currency } = this.props.company;
    return (
      <View style={styles.card} >
        <UrgencyIndicator {...this.props.twitter_data} />
        <Text>{format_name(Name)}</Text>
        <View>
          <Text>{round(LastTradePriceOnly)} {Currency}</Text>
          <Text>{PercentChange}</Text>
        </View>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  card: {
    flexDirection: 'row',
    alignSelf: 'stretch',
    paddingBottom: margin.small,
    marginHorizontal: margin.medium,
    marginTop: margin.navbar
  }
});
