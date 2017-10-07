import React, { Component } from 'react';
import { length, color } from '../../../styles/main';

const style = {
  sugestion: {
    marginBottom: length.mini,
    color: color.grey.dark
  }
}

export default class SearchSugestion extends Component {
  handleSelection = () => {
    const { name, search } = this.props;
    search(name);
  }

  render() {
    const { name } = this.props;
    return (
      <div onClick={this.handleSelection} style={style.sugestion}>
        <p>{name}</p>
      </div>
    )
  }
}
