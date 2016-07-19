'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView } from 'react-native';
import { margin } from '../styles/styles';
import { values } from 'lodash';

import StockCard from './stock-card';

//Should probably be named StockList
export default class Main extends Component {
  create_user_stock_list = (stock_data, ticker_order = []) => {
    return values(stock_data);
  }

  render() {
    const { user, stocks, navigate } = this.props,
          twitter_data = (user.twitter_data.loaded) ? user.twitter_data.data : {},
          ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2}),
          user_ticker_list = ds.cloneWithRows(stocks.data);

    return (
      <View style = {styles.container}>
        <ListView
          dataSource = {user_ticker_list}
          renderHeader = {() => <Text style={styles.header}>Watchlist</Text>}
          renderRow = {(stock_data) => (
            <StockCard {...stock_data} twitter_data={twitter_data[stock_data.Symbol]} navigate={navigate}/>
          )}
        />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center',
    marginTop: margin.navbar
  },
  header: {
    alignSelf: 'flex-start',
    fontSize: 20,
    paddingLeft: 10
  }
});
