import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { untrackStock } from '../../actions/stock-actions';
import ButtonControls from '../components/button-controls';

class ButtonControlsContainer extends Component {
  untrackStock = () => {
    const { actions, state, Ticker } = this.props;
    actions.untrackStock(Ticker, state.token)
  }

  render() {
    return <ButtonControls untrackStock={this.untrackStock}/>
  }
}

export default connect(
  state => ({
    state: {
      token: state.user.token
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      untrackStock
    }, dispatch)
  })
)(ButtonControlsContainer);
