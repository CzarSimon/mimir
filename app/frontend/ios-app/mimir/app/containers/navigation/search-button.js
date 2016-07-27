'use strict';

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { toggle_search_active } from '../../actions/search.actions';
import SearchButton from '../../components/navigation/search-button';

class SearchButtonContainer extends Component {
  toggle_search() {
    this.props.actions.toggle_search_active();
  }

  render() {
    return (
      <SearchButton
        active = {this.props.state.search.active}
        action = {this.toggle_search.bind(this)}
      />
    );
  }
}
export default connect(
  (state) => ({
    state: {
      search: state.search
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      toggle_search_active
    }, dispatch)
  })
)(SearchButtonContainer);
