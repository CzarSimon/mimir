'use strict'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { updateAndRunQuery } from '../../ducks/search'
import HistoryItem from '../components/history-item'

class HistoryItemContainer extends Component {
  handleClick = () => {
    const { text, actions } = this.props
    actions.updateAndRunQuery(text)
  }

  render() {
    return <HistoryItem text={this.props.text} handleClick={this.handleClick} />
  }
}

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    updateAndRunQuery
  }, dispatch)
})

export default connect(
  state => ({}),
  dispatch => mapDispatchToProps(dispatch)
)(HistoryItemContainer)
