'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Description from './overview/description';

export default class Overview extends Component {
  render() {
    const { description, twitter_data } = this.props;
    return (
      <View style={styles.container}>
        <Description description={description} />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center'
  }
})
