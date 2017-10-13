'use strict'
import React, { Component } from 'react';
import { Keyboard, View } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import SearchButton from '../../components/navigation/search-button';
import { getSearchRoute, SEARCH_PAGE } from '../../routing/main';
import { activateSearchKeyboard, cancelSearch } from '../../ducks/search';

class SearchButtonContainer extends Component {
  goToSearch = () => {
    const { navigator, actions, index } = this.props;
    actions.activateSearchKeyboard();
    navigator.push(getSearchRoute(index));
  }

  cancelSearch = () => {
    this.props.actions.cancelSearch();
    Keyboard.dismiss();
  }

  userLoggedIn = () => {
    const { id, token } = this.props.state.user;
    return (id && token);
  }

  render() {
    const { route, actions, state } = this.props;
    const active = route.name === SEARCH_PAGE;
    return ((this.userLoggedIn()) ?
      (
        <SearchButton
          active={active}
          goToSearch={this.goToSearch}
          cancelSearch={this.cancelSearch}
          keyboardDown={state.search.keyboardDown}
          />
      ) : (
        <View />
      )
    );
  }
}

const mapStateToProps = state => ({
  state: {
    search: state.search,
    user: state.user
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
