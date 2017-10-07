import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { browserHistory } from 'react-router';
import { connect } from 'react-redux';
import { activateSearchKeyboard, cancelSearch } from '../../../ducks/search';

import SearchButton from '../components/search-button';

class SearchButtonContainer extends Component {
  goToSearch = () => {
    this.props.actions.activateSearchKeyboard();
    browserHistory.push('/search');
  }

  cancelSearch = () => {
    browserHistory.goBack();
  }

  render() {
    const { pathname, search } = this.props.state;
    return (
      <SearchButton
        active={pathname === '/search'}
        goToSearch={this.goToSearch}
        cancelSearch={this.cancelSearch}
        keyboardDown={search.keyboardDown} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    search: state.search,
    pathname: state.routing.locationBeforeTransitions.pathname
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    activateSearchKeyboard,
    cancelSearch
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(SearchButtonContainer)
