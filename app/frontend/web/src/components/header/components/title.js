import React, { Component } from 'react';
import { color, font, length } from '../../../styles/main';

const style = {
  logo: {
    paddingTop: length.mini,
    textAlign: 'center',
    color: color.white,
    margin: '0 auto',
    fontSize: font.size.small,
    flex: 8,
    overflow: 'auto'
  }
}

export default class Title extends Component {
  render() {
    return <p style={style.logo}>{this.props.text}</p>
  }
}
