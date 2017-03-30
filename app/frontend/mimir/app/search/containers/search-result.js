'use strict'
import React, { Component } from 'react'
import { View } from 'react-native'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { addTicker } from '../../ducks/user'
import { toggleSearchActive } from '../../ducks/search'
import SearchResult from '../components/search-result'

class SearchResultContainer extends Component {
  addNewTicker = ticker => {
    const { addTicker, toggleSearchActive } = this.props.actions
    addTicker(ticker)
  }

  render() {
    const { props, addNewTicker } = this
    return (
      <SearchResult
        resultInfo={props.resultInfo}
        addTicker={addNewTicker}
        goToStock={props.goToStock} />
    )
  }
}

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    addTicker,
    toggleSearchActive
  }, dispatch)
})

export default connect(
  state => ({}),
  dispatch => mapDispatchToProps(dispatch)
)(SearchResultContainer)
