'use strict';
import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native';
import Description from './overview/description';
import TwitterDiagram from './overview/twitter-diagram';
import Topics from './overview/topics';
import { length } from '../styles/styles';

export default class Overview extends Component {
  render() {
    const { description, twitter_data } = this.props;
    return (
      <View style={styles.container}>
        <Description description={description} />
        <TwitterDiagram data={twitter_data} />
        <Topics data={twitter_data}/>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'flex-start'
  }
})
