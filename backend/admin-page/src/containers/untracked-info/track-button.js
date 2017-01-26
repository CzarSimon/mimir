import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { startTrackingTicker } from '../../actions/tickers-actions';
import TrackButton from '../../components/untracked-info/track-button';

class TrackButtonContainer extends Component {
  startTrackingTicker = () => {
    const { companyName, tickerName, description, imageUrl, website } = this.props
    console.log(this.props);
    const { token } = this.props.state
    const { startTrackingTicker } = this.props.actions
    startTrackingTicker(tickerName, companyName, description, imageUrl, website, token)
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
