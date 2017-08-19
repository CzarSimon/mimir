'use strict'

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { updateAndRunQuery, activateSearchKeyboard } from '../../ducks/search';
import { appendToSearchHistory } from '../../ducks/user';
import { SEARCH_PAGE } from '../../routing/main';
import Title from '../../components/navigation/title';
import SearchBar from '../../components/navigation/search-bar';

class TitleContainer extends Component {
  runQuery = query => {
    this.props.actions.updateAndRunQuery(query);
  }

  addQuery = query => {
    const { id } = this.props.state.user;
    this.props.actions.appendToSearchHistory(id, query);
  }

  render() {
    const { search } = this.props.state;
    const { title, name } = this.props.route;
    const { activateSearchKeyboard } = this.props.actions;
    return (
      (name !== SEARCH_PAGE)
      ? <Title title={title} />
      : (
          <SearchBar
            runQuery={this.runQuery}
            query={search.query}
            addQuery={this.addQuery}
            keyboardDown={search.keyboardDown}
            activateSearchKeyboard={activateSearchKeyboard}
          />
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
    updateAndRunQuery,
    appendToSearchHistory,
    activateSearchKeyboard
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(TitleContainer)
