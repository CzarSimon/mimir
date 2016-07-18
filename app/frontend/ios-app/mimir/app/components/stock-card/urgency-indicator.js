import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native'
import { color, margin } from '../../styles/styles';
import { urgency_level } from '../../methods/server/twitter-miner';

export default class UrgencyIndicator extends Component {
  render() {
    const { volume, mean, stdev, minute } = this.props;
    switch (urgency_level(volume, mean, stdev, minute)) {
      case 'high':
        return (<View style={styles.urgent} />);
      case 'urgent':
        return (<View style={styles.very_urgent} />);
      default:
        return (<View />);
    }
  }
}

const styles = StyleSheet.create({
  very_urgent: {
    backgroundColor: color.red,
    width: margin.medium
  },
  urgent: {
    backgroundColor: color.yellow,
    width: margin.medium
  }
});
