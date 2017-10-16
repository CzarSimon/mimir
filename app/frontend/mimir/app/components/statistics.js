'use strict'

import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Row from './statistics/row';
import PriceChart from './statistics/price-chart';
import Separator from './helpers/separator';
import { length, font } from '../styles/styles';
import { round, formatThousands } from '../methods/helper-methods';

export default class Statistics extends Component {
  render() {
    const {
      Ask, Bid, MarketCapitalization, EarningsShare,
      Volume, EBITDA, PERatio, YearHigh, YearLow,
      ChangeFromYearHigh, ChangeFromYearLow,
      Open, AverageDailyVolume
    } = this.props;
    //historicalData
    //<PriceChart historicalData={historicalData} />
    return (
      <View style={styles.container}>
        <View style={styles.titleGroup}>
          <Text style={styles.title}>Fundamentals</Text>
          <Separator />
        </View>
        <Row name={"Opening Price"} value={round(Ask)} />
        <Row name={"Ask"} value={round(Ask)} />
        <Row name={"Bid"} value={round(Bid)} />
        <Row name={"EPS"} value={round(EarningsShare)} />
        <Row name={"Volume"} value={formatThousands(Volume)} />
        <Row name={"Avg Daily Volume"} value={formatThousands(AverageDailyVolume)} />
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
    justifyContent: 'flex-start'
  },
  title: {
    fontFamily: font.type.sans.bold,
    fontSize: font.h5,
    marginBottom: length.small
  },
  titleGroup: {
    marginLeft: length.medium,
    marginVertical: length.small
  }
});
