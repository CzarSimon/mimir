'use strict';

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import socket from '../methods/server/socket';
import { fetch_news_items, recive_news_items } from '../actions/news.actions';
import Newslist from '../components/newslist';
import Loading from '../components/loading';

class NewsContainer extends Component {
  componentWillMount() {
    const { company, actions, state } = this.props;
    const { active_ticker } = state.navigation;
    actions.fetch_news_items(active_ticker, socket);
  }

  render() {
    const { navigation, news }  = this.props.state;
    const company_news = news[navigation.active_ticker];
    const component = (company_news) ? (<Newslist news={company_news} />) : (<Loading />);
    return component;
  }
}


export default connect(
  state => ({
    state: {
      news: state.news,
      navigation: state.navigation
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      fetch_news_items
    }, dispatch)
  })
)(NewsContainer);
