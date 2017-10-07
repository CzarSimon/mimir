import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { browserHistory } from 'react-router';
import { connect } from 'react-redux';
import { reciveToken } from '../../../ducks/user';
import { logonUser } from '../../../ducks/logon';
import Loading from '../../helpers/loading';

class CallbackContainer extends Component {
  storeTokenAndRedirect = token => {
    const { reciveToken, logonUser } = this.props.actions;
    reciveToken(token);
    logonUser();
    browserHistory.replace('/');
  }

  componentDidMount() {
    this.props.auth.handleAuthentication(this.storeTokenAndRedirect);
  }

  render() {
    return <Loading />
  }
}

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    reciveToken,
    logonUser
  }, dispatch)
});

export default connect(
  state => ({}),
  dispatch => mapDispatchToActions(dispatch)
)(CallbackContainer);
