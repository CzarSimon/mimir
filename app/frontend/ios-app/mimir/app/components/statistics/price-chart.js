'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { color, length, font } from '../../styles/styles'
import { card } from '../../styles/common'
import Chart from 'react-native-chart'
import { map, reverse } from 'lodash'

export default class PriceChart extends Component {
  render() {
    const { historicalData } = this.props
    const data = reverse(map(historicalData, (obj) => [obj.Date, obj.Adj_Close]))
    return (
      <View style={styles.container}>
        <Text style={styles.chartLedgend}>Price chart (3M)</Text>
        <Chart
          style={styles.chart}
          data={data}
          type='bar'
          lineWidth={3}
          showXAxisLabels={false}
          tightBounds={true}
          showGrid={false}
          color={color.blue}
          axisColor={color.black}
          axisLabelColor={color.black}
        />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: 'stretch',
    justifyContent: 'center',
    alignItems: 'stretch',
    marginVertical: length.small,
    margin: length.medium,
    ...card
  },
  chartLedgend: {
    marginTop: length.small,
    marginLeft: length.medium,
    fontFamily: font.type.sans.bold,
    fontSize: font.h4
  },
  chart: {
    padding: length.small,
    paddingVertical: length.medium,
    height: 1.8 * length.navbar
  }
})
