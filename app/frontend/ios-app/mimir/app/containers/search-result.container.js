//Start here tomorrow
'use strict';

import React, { Component } from 'react';
import { View } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import SearchResult from '../components/search-result';

class SearchResultContainer extends Component {
  render() {
    const { active, results } = this.props.state.search;
    if (active && results.length) {
      return (<SearchResult results={results}/>)
    } else {
      return (<View />);
    }
  }
}

export default connect(
  (state) => ({
    state: {
      search: state.search
    }
  })
)(SearchResultContainer)
