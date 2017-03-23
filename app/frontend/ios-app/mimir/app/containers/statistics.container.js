'use strict';
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { fetchStockData, fetchHistoricalData } from '../ducks/stocks'
import { retriveHistoricalData } from '../methods/yahoo-api'
import Statistics from '../components/statistics'
import Loading from '../components/loading'

class StatisticsContainer extends Component {
  componentWillMount() {
    const { activeTicker } = this.props.state.navigation;
    const { fetchStockData, fetchHistoricalData } = this.props.actions;
    fetchStockData([activeTicker]);
    fetchHistoricalData(activeTicker);
  }
  render() {
    const { stocks, navigation } = this.props.state;
    const company = stocks.data[navigation.activeTicker];
    if (company.EBITDA && company.historicalData) {
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
      fetchStockData,
      fetchHistoricalData
    }, dispatch)
  })
)(StatisticsContainer);
