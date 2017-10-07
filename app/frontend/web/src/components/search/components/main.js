import React, { Component } from 'react';
import SearchResults from './search-results';
import SearchSugestions from './search-sugestions';

export default class Search extends Component {
  render() {
    const { results, sugestions, addNewTicker, updateAndRunQuery, userId } = this.props;
    return (results.length) ?
      <SearchResults
        userId={userId}
        results={results}
        addNewTicker={addNewTicker} /> :
      <SearchSugestions
        sugestions={sugestions}
        updateAndRunQuery={updateAndRunQuery} />
  }
}
