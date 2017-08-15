'use strict'
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import SearchButton from '../../components/navigation/search-button';
import { getSearchRoute, SEARCH_PAGE } from '../../routing/main';

class SearchButtonContainer extends Component {
  goToSearch = () => {
    const { navigator, actions, index } = this.props;
    navigator.push(getSearchRoute(index));
  }

  render() {
    const active = this.props.route.name === SEARCH_PAGE;
    return (
      <SearchButton active={active} goToSearch={this.goToSearch} />
    );
  }
}
export default connect(
  state => ({}),
  dispatch => ({})
)(SearchButtonContainer)
