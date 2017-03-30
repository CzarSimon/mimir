'use strict'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import BackButton from '../../components/navigation/back-button'

class BackButtonContainer extends Component {
  handleClick = () => {
    const { navigator } = this.props
    navigator.pop()
  }

  render() {
    const { index } = this.props
    return <BackButton index={index} handleClick={this.handleClick} />
  }
}

export default connect(
  state => ({}),
  dispatch => ({})
)(BackButtonContainer)
