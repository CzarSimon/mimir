'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, length } from '../../styles/styles';
import Chart from 'react-native-chart';
import { map, reverse } from 'lodash';

export default class PriceChart extends Component {
  render() {
    const { historical_data } = this.props;
    const data = reverse(map(historical_data, (obj) => [obj.Date, obj.Adj_Close]))
    return (
      <View style={styles.container}>
        <Chart
          style={styles.chart}
          data={data}
          type='bar'
          lineWidth={3}
          showXAxisLabels={false}
          tightBounds={true}
          showGrid={false}
          color={color.green}
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
    marginVertical: length.mini,
    paddingLeft: length.mini,
    paddingVertical: length.medium
  },
  chart: {
    marginRight: length.medium,
    height: 1.5 * length.navbar
  }
})
