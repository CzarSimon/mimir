'use strict';

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { toggle_modifiable } from '../../actions/user.actions';
import Header from '../../components/watchlist/header';

class HeaderContainer extends Component {
  handle_click() {
    this.props.actions.toggle_modifiable();
  }

  render() {
    const { state, actions } = this.props;
    return <Header modifiable={state.user.modifiable} handle_click={this.handle_click.bind(this)}/>;
  }
}

export default connect(
  (state) => ({
    state: {
      user: state.user
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      toggle_modifiable
    }, dispatch)
  })
)(HeaderContainer)
