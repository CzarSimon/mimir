import React, { Component } from 'react';
import { length, color, font } from '../../styles/styles';

const styles = {
  button: {
    padding: length.small,
    paddingLeft: length.medium,
    paddingRight: length.medium,
    backgroundColor: color.green,
    color: color.white,
    fontSize: font.size.small,
    borderWidth: 0
  }
}

export default class TrackButton extends Component {
  handleClick = event => {
    const target = event.target
    setTimeout(() => target.blur(), 100)
    this.props.handleClick()
  }

  render() {
    return (
      <button
        style={styles.button}
        onClick={this.handleClick}>
        Track Ticker
      </button>
    )
  }
}
