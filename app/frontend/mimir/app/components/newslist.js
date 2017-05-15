'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView } from 'react-native';
import { color, length } from '../styles/styles';
import Header from './newslist/header';
import NewsCard from './newslist/news-card';
import NoNews from './newslist/no-news';

export default class Newslist extends Component {
  render() {
    if (this.props.news.length > 0) {
      const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2});
      const newsList = ds.cloneWithRows(this.props.news);
      return (
        <View style={styles.container}>
          <ListView
            dataSource = {newsList}
            renderHeader = {() => (<Header />)}
            renderRow = {(articleInfo) => (<NewsCard articleInfo={articleInfo} />)}
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
    justifyContent: 'center',
    marginBottom: length.button
  }
})
