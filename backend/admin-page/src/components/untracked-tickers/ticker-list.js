import React, { Component } from 'react';
import TickerCard from './ticker-card';

export default class TickerList extends Component {
  render() {
    const { tickers } = this.props;
    return (
      <div>
        {tickers.map((ticker, i) => (
          <TickerCard name={ticker.Name} count={ticker.Observances} key={i} />)
        )}
      </div>
    )
  }
}
