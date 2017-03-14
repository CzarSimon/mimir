'use strict'

import React, { Component } from 'react'
import { View, Text } from 'react-native'
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import SearchHistory from '../components/search-history'
import SearchResults from '../components/search-results'

class SearchContainer extends Component {
  render() {
    const { search, searchHistory } = this.props.state
    return {
      (!search.query)
      ? <SearchHistory history={searchHistory}/>
      : <SearchResults results={search.results}/>
    }
  }
}

export default connect(
  state => ({
    state: {
      search: state.search,
      searchHistory: state.user.searchHistory
    }
  })
)(SearchContainer)
