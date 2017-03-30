import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native'
import { color, length } from '../../../styles/styles';
import { urgency_level } from '../../../methods/server/twitter-miner';
import Icon from 'react-native-vector-icons/Ionicons';

export default class UrgencyIndicator extends Component {
  render() {
    const { volume, mean, stdev, minute } = this.props;
    switch (urgency_level(volume, mean, stdev, minute)) {
      case 'high':
        return (
          <View style={styles.alert}>
            <Icon name='ios-alert-outline' size={length.icons.medium} color={color.yellow} />
          </View>
        );
      case 'urgent':
        return (
          <View style={styles.alert}>
            <Icon name='ios-alert-outline' size={length.icons.medium} color={color.red} />
          </View>
        );
      default:
        return (<View />);
    }
  }
}

const styles = StyleSheet.create({
  alert: {
    paddingRight: length.small,
    justifyContent: 'center'
  }
});
