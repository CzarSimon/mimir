'use strict'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { updateAndRunQuery, reciveSearchResults } from '../../ducks/search'
import { addToSearchHistory } from '../../ducks/user'
import Title from '../../components/navigation/title'
import SearchBar from '../../components/navigation/search-bar'
import socket from '../../methods/server/socket'

class TitleContainer extends Component {
  componentDidMount() {
    socket.on('DISPATCH_SEARCH_RESULTS', data => {
      this.props.actions.reciveSearchResults(data.results)
    })
  }

  runQuery = query => {
    this.props.actions.updateAndRunQuery(query)
  }

  addQuery = query => {
    this.props.actions.addToSearchHistory(query)
  }

  render() {
    const { search } = this.props.state
    return (
      (!search.active)
      ? <Title title={this.props.title} />
      : <SearchBar runQuery={this.runQuery} query={search.query} addQuery={this.addQuery} />
    )
  }
}

export default connect(
  state => ({
    state: {
      search: state.search,
      user: state.user
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      reciveSearchResults,
      updateAndRunQuery,
      addToSearchHistory
    }, dispatch)
  })
)(TitleContainer)
