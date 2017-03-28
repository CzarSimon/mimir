'use strict'
import React, { Component } from 'react'
import { View, ScrollView, Text, StyleSheet } from 'react-native'
import Row from './statistics/row'
import PriceChart from './statistics/price-chart'
import Separator from './helpers/separator'
import { length } from '../styles/styles'
import { round, formatThousands } from '../methods/helper-methods'

export default class Statistics extends Component {
  render() {
    const {
      Ask, Bid, MarketCapitalization, EarningsShare,
      Volume, EBITDA, PERatio, YearHigh, YearLow,
      ChangeFromYearHigh, ChangeFromYearLow,
      Open, AverageDailyVolume,
      historicalData
    } = this.props
    return (
      <ScrollView>
        <View style={styles.container}>
          <PriceChart historicalData={historicalData} />
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
      </ScrollView>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'flex-start',
    marginBottom: length.button + length.medium
  }
})
