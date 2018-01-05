import React, { Component } from 'react';
import { length, font } from '../../styles/styles';

const styles = {
  text: {
    paddingBottom: length.mini,
    fontSize: font.size.small,
    lineHeight: '140%'
  }
}

export default class Description extends Component {
  render() {
    const { text, customStyles } = this.props;
    const combinedStyle = {
      ...styles.text,
      ...customStyles
    }
    return (
      <p style={combinedStyle}>{text}</p>
    )
  }
}
