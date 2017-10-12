import React, { Component } from  'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { logoutUser } from '../../../ducks/user';
import LogoutButton from '../components/logout';

class LogoutContainer extends Component {
  logout = () => {
    this.props.actions.logoutUser();
  }

  render() {
    return <LogoutButton logout={this.logout} />
  }
}

const mapStateToProps = state => ({})

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    logoutUser
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(LogoutContainer)
