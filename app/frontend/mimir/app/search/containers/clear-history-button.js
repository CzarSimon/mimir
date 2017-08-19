'use strict'

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { deleteSearchHistory } from '../../ducks/user';
import ClearHistoryButton from '../components/clear-history-button';

class ClearHistoryContainer extends Component {
  clearHistory = () => {
    const { actions, state } = this.props;
    actions.deleteSearchHistory(state.userId);
  }

  render() {
    return <ClearHistoryButton clearHistory={this.clearHistory} />
  }
}

const mapStateToProps = state => ({
  state: {
    userId: state.user.id
  }
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    deleteSearchHistory
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToProps(dispatch)
)(ClearHistoryContainer);
