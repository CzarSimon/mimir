import React, { Component } from 'react';
import { browserHistory } from 'react-router';

export default class MenuItem extends Component {
  handleClick = () => {
    browserHistory.push(this.props.path);
  }

  render() {
    const { idName, name } = this.props;
    return (
      <a id={idName} onClick={this.handleClick}>
        {name}
      </a>
    )
  }
}
