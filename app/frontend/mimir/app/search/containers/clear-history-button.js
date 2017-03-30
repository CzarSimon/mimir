'use strict'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { clearSearchHistory } from '../../ducks/user'
import ClearHistoryButton from '../components/clear-history-button'

class ClearHistoryContainer extends Component {
  clearHistory = () => {
    this.props.actions.clearSearchHistory()
  }

  render() {
    return <ClearHistoryButton clearHistory={this.clearHistory} />
  }
}

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    clearSearchHistory
  }, dispatch)
})

export default connect(
  state => ({}),
  dispatch => mapDispatchToProps(dispatch)
)(ClearHistoryContainer)
