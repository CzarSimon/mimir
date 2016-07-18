'use strict';

import React, { Component } from 'react';
import { Platform } from 'react-native'
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import io from 'socket.io-client/socket.io';

import Main from '../components/main';
import * as user_actions from '../actions/user.actions';
import * as stock_actions from '../actions/stock.actions';
import * as twitter_data_actions from '../actions/twitter_data.actions';
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

  render() {
    const { state, actions } = this.props;
    return (
      <Main
        user = {state.user}
        stocks = {state.stocks}
        socket = {this.socket}
        {...actions}
      />
    );
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
