import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { startTrackingTicker } from '../../actions/tickers-actions';
import TrackButton from '../../components/untracked-info/track-button';

class TrackButtonContainer extends Component {
  startTrackingTicker = () => {
    const { companyName, tickerName, description, actions, state } = this.props;
    actions.startTrackingTicker(tickerName, companyName, description, state.token)
  }

  render() {
    return <TrackButton handleClick={() => this.startTrackingTicker()}/>
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
      startTrackingTicker
    }, dispatch)
  })
)(TrackButtonContainer);
