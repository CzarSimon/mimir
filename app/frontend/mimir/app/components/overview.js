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
      <ScrollView style={style.scrollView}>
        <View style={style.container}>
          <Description description={description} />
          <TwitterDiagram data={twitterData} />
          <Statistics {...company} />
        </View>
      </ScrollView>
    );
  }
}

const style = StyleSheet.create({
  scrollView: {
    marginBottom: length.button
  },
  container: {
    alignItems: 'stretch',
    justifyContent: 'flex-start'
  }
})
