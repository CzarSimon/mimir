import React, { Component } from 'react';
import { color, font } from '../../../styles/main';
import SearchSugestion from './search-sugestion';

const style = {
  container: {
    textAlign: 'center'
  },
  title: {
    color: color.blue,
    fontSize: font.size.medium,
  }
}

export default class SearchSugestions extends Component {
  render() {
    const { sugestions, updateAndRunQuery } = this.props;
    return (
      <div style={style.container}>
        <p style={style.title}>Sugestions</p>
        {sugestions.map((sugestion, i) => (
          <SearchSugestion key={i} {...sugestion } search={updateAndRunQuery} />
        ))}
      </div>
    )
  }
}
