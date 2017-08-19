'use strict'

import React, { Component } from 'react';
import { View, StyleSheet, ScrollView, } from 'react-native';
import Description from './overview/description';
import TwitterDiagram from './overview/twitter-diagram';
import Topics from './overview/topics';
import Statistics from './statistics';
import { length } from '../styles/styles';

export default class Overview extends Component {
  render() {
    const { description, twitterData, company } = this.props;
    //<Topics data={twitterData}/>
    return (
      <ScrollView>
        <View style={styles.container}>
          <Description description={description} />
          <TwitterDiagram data={twitterData} />
          <Statistics {...company} />
        </View>
      </ScrollView>
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
