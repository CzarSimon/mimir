import React, { Component } from 'react';
import { length, color, font } from '../../../styles/main';

const style = {
  button: {
    color: color.grey.dark,
    width: '50%',
    margin: `${length.navbar} auto`,
    marginBottom: `-${length.large}`,
    border: 'solid',
    borderRadius: '5px'
  },
  text: {
    textAlign: 'center',
    fontSize: font.size.medium
  }
}

export default class LogoutButton extends Component {
  render() {
    return (
      <div style={style.button} className='button' onClick={this.props.logout}>
        <p style={style.text}>Logout</p>
      </div>
    )
  }
}
