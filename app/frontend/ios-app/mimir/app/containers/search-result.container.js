//Start here tomorrow
'use strict';

import React, { Component } from 'react';
import { View } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { add_ticker } from '../actions/user.actions';
import { toggle_search_active } from '../actions/search.actions';
import SearchResult from '../components/search-result';

class SearchResultContainer extends Component {
  add_new_ticker(new_ticker) {
    const { add_ticker, toggle_search_active } = this.props.actions;
    add_ticker(new_ticker);
    toggle_search_active();
  }

  render() {
    const { active, results } = this.props.state.search;
    if (active && results.length) {
      return (
        <SearchResult
          results = {results}
          add_ticker = {this.add_new_ticker.bind(this)}
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
      add_ticker,
      toggle_search_active
    }, dispatch)
  })
)(SearchResultContainer)
