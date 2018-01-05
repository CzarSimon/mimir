import React, { Component } from 'react';
import { length } from '../../styles/styles';

const styles = {
  text: {
    paddingBottom: length.mini
  }
}

export default class InfoName extends Component {
  render() {
    return (
      <h2 style={styles.text}>{this.props.name}</h2>
    )
  }
}
