'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView, FlatList } from 'react-native';
import { color, length } from '../styles/styles';
import Header from './newslist/header';
import NewsCard from './newslist/news-card';
import NoNews from './newslist/no-news';

export default class Newslist extends Component {
  render() {
    if (this.props.news.length > 0) {
      return (
        <View style={styles.container}>
          <FlatList
            data={this.props.news}
            ListHeaderComponent={() => <Header />}
            keyExtractor={(item, index) => index}
            renderItem={({ item }) => (<NewsCard articleInfo={item} />)}
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
