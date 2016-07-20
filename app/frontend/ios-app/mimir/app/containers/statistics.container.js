'use strict';
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { fetch_stock_data } from '../actions/stock.actions';
import Statistics from '../components/statistics';
import Loading from '../components/loading';

class StatisticsContainer extends Component {
  componentWillMount() {
    const { actions, state } = this.props;
    actions.fetch_stock_data([state.navigation.active_ticker]);
  }
  render() {
    const { stocks, navigation } = this.props.state;
    const company = stocks.data[navigation.active_ticker];
    if (company.EBITDA) {
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
      fetch_stock_data
    }, dispatch)
  })
)(StatisticsContainer);
