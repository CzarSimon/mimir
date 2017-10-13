import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchAndInitalizeUser } from '../../../ducks/logon';
import { logonUser } from '../../../ducks/user';
import { login, storeCredentials, parseCredentials } from '../../../methods/auth-service';
import { MAIN_ROUTE } from '../../../routing/main';
import Login from '../components/main';

class LoginContainer extends Component {
  loginWith = provider => {
    const { logonUser, fetchAndInitalizeUser } = this.props.actions;
    login(provider)
    .then(credentials => {
      parseCredentials(credentials.accessToken)
      .then(id => {
        storeCredentials({...credentials, id});
        logonUser(id, credentials.idToken);
        fetchAndInitalizeUser(id);
      })
      .catch(console.log)
    })
  }

  render() {
    return <Login loginWith={this.loginWith} />
  }
}

const mapStateToProps = state => ({})

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    fetchAndInitalizeUser,
    logonUser
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(LoginContainer)
