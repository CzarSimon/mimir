'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { color, font, length } from '../../styles/styles';
import { create_subject_string, create_clean_title } from '../../methods/helper-methods';
import Summary from './news-card/summary';
import Info from './news-card/info';
import SafariView from 'react-native-safari-view';

export default class NewsCard extends Component {
  handle_click = () => {
    SafariView.isAvailable()
    .then(SafariView.show({
      url: this.props.article_info.url,
      tintColor: color.blue
    }))
    .catch(err => {console.log(err)})
  }
  render() {
    const { title, compound_score, timestamp, twitter_references, summary } = this.props.article_info;
    const clean_title = create_clean_title(title);
    const component = (summary) ? <Summary clicked={true} summary={summary} /> :
    (
      <Info
        twitter_references={twitter_references}
        compound_score={compound_score}
        timestamp={timestamp}
      />
    )
    return (
      <TouchableHighlight
        onPress = {() => this.handle_click()}
        underlayColor = {color.grey.background}>
        <View style={styles.card}>
          <Text style={styles.title}>{clean_title}</Text>
          {component}
        </View>
      </TouchableHighlight>
    );
  }
}

const styles = StyleSheet.create({
  card: {
    flex: 1,
    alignItems: 'stretch',
    padding: length.small,
    marginHorizontal: length.medium,
    marginBottom: length.small,
    borderColor: color.grey.background,
    borderWidth: 1,
    borderBottomWidth: 2,
    backgroundColor: color.white
  },
  title: {
    fontSize: font.h4,
    fontFamily: font.type.sans.bold,
    color: color.black,
    marginBottom: length.mini
  }
})
