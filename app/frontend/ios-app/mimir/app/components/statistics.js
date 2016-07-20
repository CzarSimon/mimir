'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Row from './statistics/row';
import { length } from '../styles/styles';
import { round, format_thousands } from '../methods/helper-methods';

export default class Statistics extends Component {
  render() {
    const {
      Ask, Bid, MarketCapitalization, EarningsShare,
      Volume, EBITDA, PERatio, YearHigh, YearLow,
      ChangeFromYearHigh, ChangeFromYearLow,
      Open, PreviousClose, AverageDailyVolume
    } = this.props;
    return (
      <View style={styles.container}>
        <Row name={"Opening Price"} value={round(Ask)} />
        <Row name={"Previous Close"} value={round(Ask)} />
        <Row name={"Ask"} value={round(Ask)} />
        <Row name={"Bid"} value={round(Bid)} />
        <Row name={"EPS"} value={round(EarningsShare)} />
        <Row name={"Volume"} value={format_thousands(Volume)} />
        <Row name={"Avg Daily Volume"} value={format_thousands(AverageDailyVolume)} />
        <Row name={"Market Cap"} value={MarketCapitalization} />
        <Row name={"EBITDA"} value={EBITDA} />
        <Row name={"PE Ratio"} value={round(PERatio)} />
        <Row name={"One Year High"} value={round(YearHigh)} />
        <Row name={"% From One Year High"} value={round(ChangeFromYearHigh) + " %"} />
        <Row name={"One Year Low"} value={round(YearLow)} />
        <Row name={"% From One Year Low"} value={round(ChangeFromYearLow) + " %"} />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'flex-start',
    marginLeft: length.medium
  }
})
