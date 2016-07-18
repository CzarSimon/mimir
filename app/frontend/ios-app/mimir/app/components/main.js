'use strict';

import React, { Component } from 'react';
import {
  View, Text, StyleSheet,
  ListView,
  ActivityIndicator,
  TouchableHighlight
} from 'react-native'
import { values } from 'lodash';

import { persist_object } from './../methods/async-storage';
import { retrive_stock_data } from './../methods/yahoo-api';
import { array_equals } from '../methods/helper-methods';
import StockCard from './stock-card';

const all_tickers = ['TWTR', 'WMT', 'GS', 'GOOG', 'FB', 'NFLX', 'AAPL', 'INTC', 'AMZN', 'NKE'];

export default class Main extends Component {
  constructor(props) {
    super(props);
  }

  componentWillMount() {
    const { fetch_user, recive_twitter_data, socket } = this.props;
    fetch_user();
    socket.on('DISPATCH TWITTER DATA', (payload) => {
      if (payload.data) { recive_twitter_data(payload.data); }
    });
  }

  componentWillReceiveProps(next_props) {
    const { user, fetch_stock_data, fetch_twitter_data, socket } = this.props;
    const { user: next_user, stocks: next_stocks } = next_props;
    if (!user.loaded && next_user.tickers.length) {
      socket.on('NEW TWITTER DATA', () => {
        fetch_twitter_data(next_user, socket);
      });
    } else if (user.loaded && !array_equals(next_user.tickers, user.tickers)) {
      persist_object("user", next_user);
      socket.on('NEW TWITTER DATA', () => {
        fetch_twitter_data(next_user, socket);
      });
    }
    if (next_user.loaded && !next_user.twitter_data.loaded) {
      fetch_twitter_data(next_user, socket);
    } else if (!next_stocks.loaded && next_user.tickers.length) {
      fetch_stock_data(next_user.tickers);
    }
  }

  create_user_stock_list = (stock_data, ticker_order = []) => {
    return values(stock_data);
  }

  render() {
    const { user, stocks, add_ticker, remove_ticker } = this.props,
          twitter_data = (user.twitter_data.loaded) ? user.twitter_data.data : {},
          ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2}),
          user_ticker_list = ds.cloneWithRows(stocks.data);

    const MainComponent = (user.loaded && stocks.loaded) ? (
      <ListView
        dataSource = {user_ticker_list}
        renderHeader = {() => <Text style={styles.header}>Watchlist</Text>}
        renderRow = {(stock_data) => (
          <StockCard {...stock_data} twitter_data={twitter_data[stock_data.Symbol]} />
        )}
      />
    ) : ( <ActivityIndicator size="large" /> );

    return (
      <View style = {styles.container}>
        { MainComponent }
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center',
    marginTop: 80
  },
  header: {
    alignSelf: 'flex-start',
    fontSize: 20,
    paddingLeft: 10
  }
});
