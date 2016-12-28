'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, length, font } from '../../styles/styles';
import Chart from 'react-native-chart';
import { map, reverse } from 'lodash';

export default class PriceChart extends Component {
  render() {
    const { historical_data } = this.props;
    const data = reverse(map(historical_data, (obj) => [obj.Date, obj.Adj_Close]))
    return (
      <View style={styles.container}>
        <Text style={styles.chart_ledgend}>Price chart (3M)</Text>
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
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: 'stretch',
    justifyContent: 'center',
    alignItems: 'stretch',
    backgroundColor: color.white,
    marginVertical: length.small,
    margin: length.medium,
    borderColor: color.grey.background,
    borderWidth: 1
  },
  chart_ledgend: {
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
