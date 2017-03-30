'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { length } from '../../styles/styles'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { addTicker } from '../../ducks/user'
import { toggleSearchActive } from '../../ducks/search'
import { companyPageRoute, getRouteIndex } from '../../routing/main'
import SearchHistory from '../components/search-history'
import SearchResults from '../components/search-results'

class SearchContainer extends Component {
  goToStock = (ticker, added = true) => {
    const { navigator } = this.props
    navigator.pop()
    /*
    const routeIndex = getRouteIndex(navigator)
    if (added) {
      navigator.replace(companyPageRoute(ticker, routeIndex))
    } else {
      navigator.push(companyPageRoute(ticker, routeIndex + 1))
    }
    */
  }

  render() {
    const { search, searchHistory } = this.props.state
    const { addTicker } = this.props.actions
    return (
      <View style={styles.container}>
        {
          (!search.query)
          ? <SearchHistory history={searchHistory} />
          : <SearchResults results={search.results} goToStock={this.goToStock}/>
        }
      </View>
    )
  }
}

const mapStateToProps = state => ({
  state: {
    search: state.search,
    searchHistory: state.user.searchHistory
  }
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    addTicker,
    toggleSearchActive
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
