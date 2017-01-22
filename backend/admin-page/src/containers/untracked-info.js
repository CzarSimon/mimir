import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchCompanyInfo } from '../actions/tickers-actions';
import UntrackedInfo from '../components/untracked-info';

class UntrackedInfoContainer extends Component {
  componentDidMount() {
    const { tickerName } = this.props.params;
    this.props.actions.fetchCompanyInfo(tickerName)
  }

  render() {
    const { tickerName } = this.props.params;
    const data = this.props.state.tickers[tickerName]
    return (
      <UntrackedInfo
        companyName={data.companyName}
        tickerName={data.Name}
        description={(data.description) ? data.description : ""}
      />
    )
  }
}

export default connect(
  state => ({
    state: {
      tickers: state.tickers.data
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      fetchCompanyInfo
    }, dispatch)
  })
)(UntrackedInfoContainer);
