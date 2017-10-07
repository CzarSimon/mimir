import React, { Component } from 'react';
import { color, length } from '../../../styles/main';
import { formatDate } from '../../../methods/helper-methods';

const style = {
  container: {
    overflow: 'auto',
    paddingTop: length.mini,
    color: color.grey.dark
  },
  references: {
    width: '50%',
    paddingLeft: length.small,
    clear: 'both',
    float: 'left'
  },
  date: {
    float: 'right',
    width: '30%',
    textAlign: 'right',
    paddingRight: length.small
  }
}

export default class Info extends Component {
  render() {
    const { twitterReferences, timestamp } = this.props;
    return (
      <div style={style.container}>
        <p style={style.references}>Tweet References: {twitterReferences.length}</p>
        <p style={style.date}>{formatDate(timestamp)}</p>
      </div>
    );
  }
}
