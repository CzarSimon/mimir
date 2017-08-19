'use strict'

import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { length } from '../../styles/styles';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { addTicker } from '../../ducks/user';
import { updateAndRunQuery, fetchSearchSugestions } from '../../ducks/search';
import { companyPageRoute, getRouteIndex } from '../../routing/main';
import SearchHistory from '../components/search-history';
import SearchSugestions from '../components/search-sugestions';
import SearchResults from '../components/search-results';

class SearchContainer extends Component {
  componentDidMount() {
    const { actions, state } = this.props;
    actions.fetchSearchSugestions(state.tickers);
  }

  goToStock = (ticker, added = true) => {
    const { navigator } = this.props;
    //navigator.pop()
    const routeIndex = getRouteIndex(navigator);
    if (added) {
      navigator.replace(companyPageRoute(ticker, routeIndex));
    } else {
      navigator.push(companyPageRoute(ticker, routeIndex + 1));
    }
  }

  render() {
    const { search, searchHistory } = this.props.state
    const { addTicker, updateAndRunQuery } = this.props.actions
    return (
      <View style={styles.container}>
        {
          (!search.query)
          ? (!search.keyboardDown)
            ? <SearchHistory history={searchHistory} />
            : <SearchSugestions sugestions={search.sugestions} updateAndRunQuery={updateAndRunQuery} />
          : <SearchResults results={search.results} goToStock={this.goToStock}/>
        }
      </View>
    )
  }
}

const mapStateToProps = state => ({
  state: {
    tickers: state.user.tickers,
    search: state.search,
    searchHistory: state.user.searchHistory
  }
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    addTicker,
    updateAndRunQuery,
    fetchSearchSugestions
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToProps(dispatch)
)(SearchContainer)

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    marginLeft: length.medium,
    marginTop: length.navbar
  }
})
