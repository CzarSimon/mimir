'use strict';

import React, { Component } from 'react';
import { Platform } from 'react-native'
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import io from 'socket.io-client/socket.io';

import Main from '../components/main';
import Loading from '../components/loading';

import * as user_actions from '../actions/user.actions';
import * as stock_actions from '../actions/stock.actions';
import * as twitter_data_actions from '../actions/twitter_data.actions';

import { persist_object } from './../methods/async-storage';
import { retrive_stock_data } from './../methods/yahoo-api';
import { array_equals } from '../methods/helper-methods';
import { SERVER_URL } from '../credentials/server-info';

class MimirApp extends Component {
  constructor(props) {
    super(props);
    this.socket = io.connect(SERVER_URL, {jsonp: false});
    this.socket.on('get info from client', (data) => {
      if (data === 'GET INFO') {
        this.socket.emit('send info to server', {
          clientMachine: (Platform.OS + " running React Native")
        });
      }
    })
  }

  componentWillMount() {
    const { fetch_user, recive_twitter_data } = this.props.actions;
    fetch_user();
    this.socket.on('DISPATCH TWITTER DATA', (payload) => {
      if (payload.data) { recive_twitter_data(payload.data); }
    });
  }

  //Well this is obviously terrible, change soon
  componentWillReceiveProps(next_props) {
    const { fetch_stock_data, fetch_twitter_data } = this.props.actions;
    const { user: next_user, stocks: next_stocks } = next_props.state;
    const { user } = this.props.state;
    const { socket } = this;

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

  render() {
    const { actions } = this.props;
    const { user, stocks } = this.props.state;

    if (user.loaded && stocks.loaded) {
      return (<Main user={user} stocks={stocks} socket={this.socket} {...actions} />);
    } else {
      return (<Loading />);
    }
  }
}

export default connect(
  (state) => ({
    state: {
      user: state.user,
      stocks: state.stocks
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      ...user_actions,
      ...stock_actions,
      ...twitter_data_actions
    }, dispatch)
  })
)(MimirApp);
