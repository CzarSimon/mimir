import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { loginUser } from '../actions/user-actions';
import Login from '../login/components/login';

class LoginContainer extends Component {
  loginSubmit = (username, password) => {
    const { loginUser } = this.props.actions
    loginUser(username, password)
  }

  render() {
    return <Login loginSubmit={this.loginSubmit} />
  }
}

export default connect(
  state => ({
    state: {
      user: state.user
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      loginUser
    }, dispatch)
  })
)(LoginContainer);
