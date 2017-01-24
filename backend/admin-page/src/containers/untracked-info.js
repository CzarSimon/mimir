import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { fetchCompanyInfo } from '../actions/tickers-actions';
import UntrackedInfo from '../components/untracked-info';

class UntrackedInfoContainer extends Component {
  componentDidMount() {
    const { tickerName } = this.props.params;
    if (this.props.state.tickersLoaded) {
      this.props.actions.fetchCompanyInfo(tickerName)
    } else {
      browserHistory.push('/untracked-tickers');
    }

  }

  render() {
    if (this.props.state.tickersLoaded) {
      const { tickerName } = this.props.params;
      const data = this.props.state.tickers[tickerName]
      return (
        <UntrackedInfo
          companyName={data.companyName}
          tickerName={data.Name}
          description={(data.description) ? data.description : ""}
          />
      )
    } else {
      return (<p>Ticker not loaded</p>)
    }
  }
}

export default connect(
  state => ({
    state: {
      tickers: state.tickers.data,
      tickersLoaded: state.tickers.loaded
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      fetchCompanyInfo
    }, dispatch)
  })
)(UntrackedInfoContainer);
