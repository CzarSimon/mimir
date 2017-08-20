'use strict'

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { color, font, length } from '../../styles/styles';
import { card } from '../../styles/common';
import { createCleanTitle } from '../../methods/helper-methods';
import ArticleSummary from './news-card/summary';
import Info from './news-card/info';
import ArticleButton from './news-card/article-button';
import SafariView from 'react-native-safari-view';

export default class NewsCard extends Component {
  handleClick = () => {
    const { url } = this.props.articleInfo;
    SafariView.isAvailable()
    .then(SafariView.show({
      url,
      tintColor: color.blue
    }));
  }

  summaryComponent = () => {
    const { summary, url } = this.props.articleInfo;
    return (<ArticleSummary url={url} summary={summary} />);
  }

  render() {
    const { title, twitterReferences, timestamp } = this.props.articleInfo;
    const cleanTitle = createCleanTitle(title);
    return (
      <TouchableHighlight
        onPress={this.handleClick}
        underlayColor={color.grey.background}>
        <View style={styles.card}>
          <Text style={styles.title}>{cleanTitle}</Text>
          <Info
            twitterReferences={twitterReferences}
            timestamp={timestamp}
          />
        </View>
      </TouchableHighlight>
    )
  }
}

const styles = StyleSheet.create({
  card: {
    flex: 1,
    alignItems: 'stretch',
    padding: length.small,
    marginHorizontal: length.medium,
    marginBottom: length.small,
    ...card
  },
  title: {
    fontSize: font.h4,
    fontFamily: font.type.sans.bold,
    color: color.black,
    marginBottom: length.mini
  }
})
