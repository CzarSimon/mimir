'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, ListView } from 'react-native';
import { length } from '../styles/styles';
import { values } from 'lodash';

import HeaderContainer from '../containers/watchlist/header.container';
import StockCard from './watchlist/stock-card';
import SearchResultContainer from '../containers/search-result.container';

export default class Watchlist extends Component {
  create_user_stock_list = (stock_data, ticker_order = []) => {
    return values(stock_data);
  }

  render() {
    const { user, stocks, navigate, remove_ticker } = this.props,
          twitter_data = (user.twitter_data.loaded) ? user.twitter_data.data : {},
          ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2}),
          user_ticker_list = ds.cloneWithRows(stocks.data);

    return (
      <View style = {styles.container}>
        <View style={styles.search_result}>
          <SearchResultContainer />
        </View>
        <ListView
          dataSource = {user_ticker_list}
          renderHeader = {() => <HeaderContainer />}
          renderRow = {(stock_data) => {
            if (user.tickers.includes(stock_data.Symbol)) {
              return (
                <StockCard
                  {...stock_data}
                  twitter_data={twitter_data[stock_data.Symbol]}
                  navigate={navigate}
                  modifiable={user.modifiable}
                  remove_ticker={remove_ticker}
                  />
              );
            } else {
              return (<View />);
            }
          }}
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
    marginTop: length.navbar
  },
  header: {
    alignSelf: 'flex-start',
    fontSize: 20,
    paddingLeft: 10
  },
  search_result: {
    marginLeft: length.medium
  }
});
