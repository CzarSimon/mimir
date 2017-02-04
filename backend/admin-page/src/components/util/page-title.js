import React, { Component } from 'react';
import { length } from '../../styles/styles';

const styles = {
  text: {
    paddingTop: length.medium,
    paddingBottom: length.small
  }
}

export default class PageTitle extends Component {
  render() {
    const { customStyle, title} = this.props;
    const styling = {
      ...styles.text,
      ...customStyle
    }
    return (<h1 style={styling}>{title}</h1>)
  }
}
