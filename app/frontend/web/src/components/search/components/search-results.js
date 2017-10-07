import React, { Component } from 'react';
import { font, color } from '../../../styles/main';
import SearchResult from './search-result';

const style = {
  title: {
    color: color.blue,
    fontSize: font.size.medium,
  }
}

export default class SearchResults extends Component {
  render() {
    const { results, addNewTicker, userId } = this.props;
    return (
      <div>
        <p style={style.title}>Results</p>
        {results.map((results, i) => (
          <SearchResult
            key={i}
            {...results}
            userId={userId}
            addNewTicker={addNewTicker} />
        ))}
      </div>
    )
  }
}
