'use strict';
import React, { Component } from 'react'
import { View } from 'react-native'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { addTicker } from '../ducks/user'
import { toggleSearchActive } from '../ducks/search'
import SearchResult from '../components/search-result'

class SearchResultContainer extends Component {
  addNewTicker(newTicker) {
    const { addTicker, toggleSearchActive } = this.props.actions;
    addTicker(newTicker);
    toggleSearchActive();
  }

  render() {
    const { active, results } = this.props.state.search;
    if (active && results.length) {
      return (
        <SearchResult
          results = {results}
          addTicker = {this.addNewTicker.bind(this)}
        />
      )
    } else {
      return (<View />);
    }
  }
}

export default connect(
  (state) => ({
    state: {
      search: state.search,
      user: state.user
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      addTicker,
      toggleSearchActive
    }, dispatch)
  })
)(SearchResultContainer)
