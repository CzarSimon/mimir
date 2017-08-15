'use strict'

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import socket from '../methods/server/socket';
import { fetchNewsItems } from '../ducks/news';
import Newslist from '../components/newslist';
import Loading from '../components/loading';

class NewsContainer extends Component {
  componentDidMount() {
    const { company, actions, state } = this.props;
    const { activeTicker } = state.navigation;
    actions.fetchNewsItems(activeTicker, '1M');
  }

  render() {
    const { navigation, news }  = this.props.state;
    const companyNews = news[navigation.activeTicker];
    const component = (companyNews) ? (<Newslist news={companyNews} />) : (<Loading />)
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
      fetchNewsItems
    }, dispatch)
  })
)(NewsContainer)
