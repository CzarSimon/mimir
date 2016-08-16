'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView } from 'react-native';
import Header from './newslist/header';
import NewsCard from './newslist/news-card';
import NoNews from './newslist/no-news';

export default class Newslist extends Component {
  render() {
    if (this.props.news.length > 0) {
      const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2})
      , news_list = ds.cloneWithRows(this.props.news);
      return (
        <View style={styles.container}>
          <ListView
            dataSource = {news_list}
            renderHeader = {() => (<Header />)}
            renderRow = {(article_info) => (<NewsCard article_info={article_info} />)}
            />
        </View>
      );
    } else {
      return <NoNews />;
    }
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center'
  }
})
