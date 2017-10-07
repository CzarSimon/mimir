import React, { Component } from 'react';
import Spinner from 'react-spinkit';
import { color, length } from '../../styles/main';

const style = {
  container: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    marginTop: length.navbar
  }
}

export default class Loading extends Component {
  render() {
    return (
      <div style={style.container}>
        <Spinner name="line-scale" color={color.blue} />
      </div>
    )
  }
}
