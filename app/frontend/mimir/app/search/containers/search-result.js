'use strict'
import React, { Component } from 'react';
import { View } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { addNewTicker } from '../../ducks/user';
import { toggleSearchActive } from '../../ducks/search';
import SearchResult from '../components/search-result';

class SearchResultContainer extends Component {
  addNewTicker = ticker => {
    const { actions, state } = this.props;
    actions.addNewTicker(state.userId, ticker);
  }

  render() {
    const { resultInfo, goToStock } = this.props;
    return (
      <SearchResult
        resultInfo={resultInfo}
        addTicker={this.addNewTicker}
        goToStock={goToStock} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    userId: state.user.id
  }
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    addNewTicker,
    toggleSearchActive
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToProps(dispatch)
)(SearchResultContainer)
