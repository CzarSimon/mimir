import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { startTrackingTicker } from '../../actions/tickers-actions';
import Button from '../../components/util/button';
import { color } from '../../styles/styles';

class TrackButtonContainer extends Component {
  startTrackingTicker = () => {
    const { companyName, tickerName, description, imageUrl, website } = this.props
    const { token } = this.props.state
    const { startTrackingTicker } = this.props.actions
    startTrackingTicker(tickerName, companyName, description, imageUrl, website, token)
  }

  render() {
    return (
      <Button
        handleClick={() => this.startTrackingTicker()}
        customStyles={{ backgroundColor: color.green }}
        text={'Track ticker'}
      />
    )
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
