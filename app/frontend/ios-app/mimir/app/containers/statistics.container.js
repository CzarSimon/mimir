'use strict';
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { fetch_stock_data, fetch_historical_data } from '../actions/stock.actions';
import { retrive_historical_data } from '../methods/yahoo-api';
import Statistics from '../components/statistics';
import Loading from '../components/loading';

class StatisticsContainer extends Component {
  componentWillMount() {
    const { active_ticker } = this.props.state.navigation;
    const { fetch_stock_data, fetch_historical_data } = this.props.actions;
    fetch_stock_data([active_ticker]);
    fetch_historical_data(active_ticker);
  }
  render() {
    const { stocks, navigation } = this.props.state;
    const company = stocks.data[navigation.active_ticker];
    if (company.EBITDA && company.historical_data) {
      return <Statistics {...company} />
    } else {
      return <Loading />
    }
  }
}

export default connect(
  (state) => ({
    state: {
      stocks: state.stocks,
      navigation: state.navigation
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      fetch_stock_data,
      fetch_historical_data
    }, dispatch)
  })
)(StatisticsContainer);
