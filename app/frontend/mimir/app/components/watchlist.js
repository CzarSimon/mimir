'use strict'

import React, { Component } from 'react';
import { View, Text, StyleSheet, FlatList } from 'react-native';
import { length, font, color } from '../styles/styles';
import { values } from 'lodash';

import Header from './watchlist/header';
import StockCard from './watchlist/stock-card';

export default class Watchlist extends Component {
  _renderItem = ({ item }) => {
    const { user, twitterData, navigate, removeTicker } = this.props;
    return (
      <StockCard
        {...item}
        twitterData={twitterData.data[item.Symbol]}
        navigate={navigate}
        modifiable={user.modifiable}
        removeTicker={removeTicker}
      />
    );
  }

  render() {
    return (
      <View style = {styles.container}>
        <FlatList
          data={values(this.props.stocks.data)}
          ListHeaderComponent={() => <Header />}
          keyExtractor={item => item.Symbol}
          renderItem={this._renderItem}
          style={styles.list}
        />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center',
    marginTop: length.navbar
  },
  list: {
    flex: 1
  },
  header: {
    alignSelf: 'flex-start',
    fontSize: font.h3,
    paddingLeft: length.small,
  },
  search_result: {
    marginLeft: length.medium
  }
})
