import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchUntrackedTickers } from '../actions/tickers-actions';

import UntrackedTickers from '../components/untracked-tickers';

class UntrackedTickersContainer extends Component {
  componentDidMount() {
    this.props.actions.fetchUntrackedTickers();
  }

  render() {
    const { data, loaded } = this.props.state.tickers;
    return (
      <UntrackedTickers tickers={data} loaded={loaded}/>
    )
  }
}

export default connect(
  state => ({
    state: {
      tickers: state.tickers
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      fetchUntrackedTickers
    }, dispatch)
  })
)(UntrackedTickersContainer);
