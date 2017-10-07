import React, { Component } from 'react';
import { color, font, length } from '../../../styles/main';

const style = {
  input: {
    paddingTop: length.mini,
    paddingLeft: length.small,
    paddingRight: length.small,
    marginLeft: length.small,
    marginRight: length.small,
    backgroundColor: color.blue,
    fontSize: font.size.small,
    color: color.white,
    border: 'none',
    borderBottom: 'solid',
    borderBottomWidth: '1px',
    outline: 'none',
    flex: 8
  }
}

export default class SearchBar extends Component {
  handleQuery = event => {
    const { search } = this.props;
    search(event.target.value);
  }

  handleSubmit = event => {
    alert(event.target.value);
  }

  render() {
    const { query, keyboardDown } = this.props;
    return (
      <input
        className="search-bar"
        autoFocus={!keyboardDown}
        type='search'
        style={style.input}
        value={query}
        onChange={this.handleQuery}
        onSubmit={this.handleSubmit}
        placeholder='Search stocks' />
    )
  }
}
