'use strict'

import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { toggleModifiable } from '../../ducks/user'
import Header from '../../components/watchlist/header'

class HeaderContainer extends Component {
  handleClick = () => {
    this.props.actions.toggleModifiable()
  }

  render() {
    const { state, actions } = this.props
    return <Header modifiable={state.user.modifiable} handleClick={this.handleClick}/>
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
      toggleModifiable
    }, dispatch)
  })
)(HeaderContainer)
