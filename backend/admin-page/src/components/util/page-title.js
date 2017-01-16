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
    return (<h1 style={styles.text}>{this.props.title}</h1>)
  }
}
