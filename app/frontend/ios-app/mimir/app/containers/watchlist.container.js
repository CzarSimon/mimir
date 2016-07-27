'use strict';

import React, { Component } from 'react';
import { Platform } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import io from 'socket.io-client/socket.io'; //Remove

import Watchlist from '../components/watchlist';
import Loading from '../components/loading';

import * as user_actions from '../actions/user.actions';
import * as stock_actions from '../actions/stock.actions';
import * as twitter_data_actions from '../actions/twitter_data.actions';
import { set_active_ticker } from '../actions/navigation.actions';
import { logon_user } from '../actions/logon.actions';

import socket from '../methods/server/socket';

import { persist_object } from './../methods/async-storage';
import { array_equals } from '../methods/helper-methods';
import { company_page_route } from '../routing/routes';
import { SERVER_URL } from '../credentials/server-info';

class WatchlistContainer extends Component {
  constructor(props) {
    super(props);
    this.socket = socket;
  }

  componentWillMount() {
    const { logon_user, recive_twitter_data, update_stock_data } = this.props.actions;
    logon_user(this.socket);
    this.socket.on('DISPATCH TWITTER DATA', (payload) => {
      if (payload.data) { recive_twitter_data(payload.data); }
    });
    setInterval(() => {
      update_stock_data(this.props.state.user.tickers);
    }, 30000);
  }

  componentWillReceiveProps(next_props) {
    const { fetch_twitter_data, fetch_stock_data } = this.props.actions;
    const { user: next_user, stocks: next_stocks } = next_props.state;
    const { user } = this.props.state;
    const { socket } = this;

    if (!user.loaded && next_user.tickers.length) {
      socket.on('NEW TWITTER DATA', () => {
        fetch_twitter_data(next_user, socket);
      });
    } else if (user.loaded && !array_equals(next_user.tickers, user.tickers)) {
      const { id, tickers } = next_user;
      console.log(tickers);
      persist_object("user", { id, tickers });
      socket.removeListener('NEW TWITTER DATA');
      socket.on('NEW TWITTER DATA', () => {
        fetch_twitter_data(next_user, socket);
      });
      fetch_twitter_data(next_user, socket);
      fetch_stock_data(tickers);
    }
  }

  navigate_to_company(ticker) {
    const { navigator, actions } = this.props;
    actions.set_active_ticker(ticker);
    navigator.push(company_page_route(ticker));
  }

  remove_ticker(ticker) {
    this.props.actions.remove_ticker(ticker);
  }

  render() {
    const { user, stocks } = this.props.state;

    if (user.loaded && stocks.loaded) {
      return (
        <Watchlist
          user={user}
          stocks={stocks}
          remove_ticker={this.remove_ticker.bind(this)}
          navigate={this.navigate_to_company.bind(this)}
        />
      );
    } else {
      return (<Loading />);
    }
  }
}

export default connect(
  (state) => ({
    state: {
      user: state.user,
      stocks: state.stocks,
      navigation: state.navigation
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      ...user_actions,
      ...stock_actions,
      ...twitter_data_actions,
      set_active_ticker,
      logon_user
    }, dispatch)
  })
)(WatchlistContainer);
