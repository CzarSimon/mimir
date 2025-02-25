'use strict'

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchNewsItems } from '../ducks/news';
import Newslist from '../components/newslist';
import Loading from '../components/loading';

class NewsContainer extends Component {
  componentDidMount() {
    const { company, actions, state } = this.props;
    const { activeTicker } = state.navigation;
    actions.fetchNewsItems(activeTicker, state.news.period);
  }

  render() {
    const { navigation, news }  = this.props.state;
    const companyNews = news[navigation.activeTicker];
    return (companyNews) ? (<Newslist news={companyNews} />) : (<Loading />)
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
