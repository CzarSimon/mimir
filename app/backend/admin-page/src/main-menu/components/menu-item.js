import React, { Component } from 'react';
import { length, font } from '../../styles/styles';

const styles = {
  listItem: {
    paddingTop: length.medium,
    paddingBottom: length.medium,
    fontSize: font.size.medium
  }
}

export default class MenuItem extends Component {
  render() {
    const { idName, name, handleClick, selected } = this.props;
    const styling = (!selected)
    ? styles.listItem
    : {
      ...styles.listItem,
      backgroundColor: 'rgba(0,0,0,0.1)'
    }
    return (
      <li style={styling}>
        <a id={idName} onClick={handleClick}>{name}</a>
      </li>
    )
  }
}
