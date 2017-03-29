'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { length } from '../../styles/styles'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { addTicker } from '../../ducks/user'
import SearchHistory from '../components/search-history'
import SearchResults from '../components/search-results'

class SearchContainer extends Component {
  render() {
    const { search, searchHistory } = this.props.state
    const { addTicker } = this.props.actions
    return (
      <View style={styles.container}>
        {
          (!search.query)
          ? <SearchHistory history={searchHistory} />
          : <SearchResults results={search.results} addTicker={addTicker} />
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
    addTicker
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
