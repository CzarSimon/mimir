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
  handleClick = event => {
    setTimeout(() => event.target.blur(), 150)
    this.props.handleClick()
  }

  render() {
    const { text, customStyles } = this.props;
    const styling = {
      ...styles.button,
      ...customStyles
    }
    return (
      <button style={styling} onClick={this.handleClick}>
        {this.props.text}
      </button>
    )
  }
}
