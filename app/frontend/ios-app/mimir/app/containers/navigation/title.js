'use strict';

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { fetch_search_results, recive_search_results } from '../../actions/search.actions';
import Title from '../../components/navigation/title';
import SearchBar from '../../components/navigation/search-bar';
import socket from '../../methods/server/socket';

class TitleContainer extends Component {
  componentDidMount() {
    socket.on('DISPATCH_SEARCH_RESULTS', data => {
      this.props.actions.recive_search_results(data.results);
    })
  }

  run_query(query) {
    this.props.actions.fetch_search_results(query);
  }

  render() {
    const { search } = this.props.state;
    return (
      (!search.active)
      ? ( <Title title={this.props.title} /> )
      : ( <SearchBar run_query={this.run_query.bind(this)}/> )
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
      fetch_search_results,
      recive_search_results
    }, dispatch)
  })
)(TitleContainer);
