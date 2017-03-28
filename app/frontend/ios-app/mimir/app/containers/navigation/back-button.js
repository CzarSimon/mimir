'use strict'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { toggleSearchActive } from '../../ducks/search'
import { last } from 'lodash'
import { SEARCH_PAGE } from '../../routing/main'
import BackButton from '../../components/navigation/back-button'

class BackButtonContainer extends Component {
  handleClick = () => {
    const { navigator, actions } = this.props
    const lastRoute = last(navigator.getCurrentRoutes())
    if (lastRoute.name === SEARCH_PAGE) {
      actions.toggleSearchActive()
    }
    navigator.pop()
  }

  render() {
    const { index } = this.props
    return <BackButton index={index} handleClick={this.handleClick} />
  }
}

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    toggleSearchActive
  }, dispatch)
})

export default connect(
  state => ({}),
  dispatch => mapDispatchToProps(dispatch)
)(BackButtonContainer)
