'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { round } from '../../methods/helper-methods';
import { color, length, font } from '../../styles/styles';

export default class TwitterDiagram extends Component {
  constructor(props) {
    super(props);
    this.state = {
      show_vol_change: false
    }
  }

  componentWillReceiveProps(next_props) {
    if (next_props.data.volume !== this.props.data.volume) {
      this.setState({
        show_vol_change: true
      });
      setTimeout(() => {
        this.setState({
          show_vol_change: false
        });
      }, 500);
    }
  }

  render() {
    const { volume, mean, stdev, minute } = this.props.data;
    const damping = parseFloat(minute) / 60.0;
    const volume_style = (this.state.show_vol_change) ? styles.volume_change : styles.text
    return (
      <View style={styles.container}>
        <Text style={volume_style}>Tweets in the last hour: {volume}</Text>
        <Text style={styles.text}>Mean: {round(damping * mean, 0)} | Standard Deviation: {round(damping * stdev, 0)}</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    marginVertical: length.small
  },
  volume_change: {
    color: color.red,
    fontFamily: font.type.sans.bold,
    fontSize: font.text
  },
  text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  }
})
