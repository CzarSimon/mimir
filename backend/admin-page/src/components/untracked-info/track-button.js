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
  render() {
    return (
      <button
        style={styles.button}
        onClick={this.props.handleClick}>
        Track Ticker
      </button>
    )
  }
}
