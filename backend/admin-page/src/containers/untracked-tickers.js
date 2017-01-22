import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchUntrackedTickers } from '../actions/tickers-actions';
import _ from 'lodash';
import UntrackedTickers from '../components/untracked-tickers';

class UntrackedTickersContainer extends Component {
  componentDidMount() {
    this.props.actions.fetchUntrackedTickers();
  }

  render() {
    const { data, loaded } = this.props.state.tickers;
    return (
      <UntrackedTickers tickers={_.values(data)} loaded={loaded}/>
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
