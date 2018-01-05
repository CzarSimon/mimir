import React, { Component } from 'react';
import { length, color, font } from '../../styles/styles';

const styles = {
  button: {
    padding: length.small,
    paddingLeft: length.medium,
    paddingRight: length.medium,
    backgroundColor: color.blue,
    color: color.white,
    fontSize: font.size.small,
    borderWidth: 0
  }
}

export default class Button extends Component {
  handleMouseUp = event => {
    const target = event.target
    setTimeout(() => target.blur(), 100)
  }

  render() {
    const { text, customStyles, handleClick } = this.props;
    const styling = {
      ...styles.button,
      ...customStyles
    }
    return (
      <button
        style={styling}
        onClick={handleClick}
        onMouseUp={this.handleMouseUp}>
        {text}
      </button>)
  }
}
