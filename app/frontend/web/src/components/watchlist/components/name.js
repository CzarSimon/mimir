import React, { Component } from 'react';
import { length, font, color } from '../../../styles/main';
//import { formatName } from '../../../methods/helper-methods';

const style = {
  container: {
    float: 'left',
    paddingLeft: length.medium
  },
  name: {
    fontSize: font.size.medium,
    color: color.black,
    fontWeight: 'bold',
    marginBottom: length.small
  },
  ticker: {
    fontSize: font.size.normal,
    color: color.grey.dark,
    marginTop: '0'
  }
}

export default class Name extends Component {
  render() {
    const { name, ticker } = this.props;
    return (
      <div style={style.container}>
        <p style={style.name}>{name}</p>
        <p style={style.ticker}>{ticker}</p>
      </div>
    )
  }
}
