'use strict'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { fetchSearchResults, reciveSearchResults } from '../../ducks/search'
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
    this.props.actions.fetchSearchResults(query)
  }

  render() {
    const { search } = this.props.state
    return (
      (!search.active)
      ? ( <Title title={this.props.title} /> )
      : ( <SearchBar run_query={this.runQuery}/> )
    )
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
      fetchSearchResults,
      reciveSearchResults
    }, dispatch)
  })
)(TitleContainer)
