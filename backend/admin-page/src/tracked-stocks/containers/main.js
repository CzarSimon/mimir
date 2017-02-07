import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchTrackedStocks } from '../../actions/stock-actions';
import TrackedStocks from '../components/main'

class TrackedStocksContainer extends Component {
  componentDidMount() {
    const { actions, state } = this.props;
    actions.fetchTrackedStocks(state.user.token)
  }

  render() {
    return (<TrackedStocks {...this.props.state.stocks}/>)
  }
}

export default connect(
  state => ({
    state: {
      user: state.user,
      stocks: state.stocks
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      fetchTrackedStocks
    }, dispatch)
  })
)(TrackedStocksContainer);
