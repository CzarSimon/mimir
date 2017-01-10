import React, { Component } from 'react';
import { browserHistory } from 'react-router';

export default class TickerCard extends Component {
  handleClick = () => {
    browserHistory.push(`/ticker/${this.props.name}`);
  }

  render() {
    return (
      <div className="ticker-card" onClick={this.handleClick}>
        <p>{this.props.name}</p>
      </div>
    )
  }
}
