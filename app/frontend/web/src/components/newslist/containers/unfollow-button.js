import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { browserHistory } from 'react-router';
import { connect } from 'react-redux';
import { deleteTicker } from '../../../ducks/user';
import UnfollowButton from '../components/unfollow-button';

class UnfollowButtonContainer extends Component {
  unfollowTicker = () => {
    const { activeTicker, userId }  = this.props.state;
    this.props.actions.deleteTicker(userId, activeTicker);
    browserHistory.goBack();
  }

  render() {
    return (
      <UnfollowButton unfollowTicker={this.unfollowTicker} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    userId: state.user.id,
    activeTicker: state.navigation.activeTicker
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    deleteTicker
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(UnfollowButtonContainer);
