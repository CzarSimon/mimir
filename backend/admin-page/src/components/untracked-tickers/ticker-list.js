import React, { Component } from 'react';
import TickerCard from './ticker-card';

export default class TickerList extends Component {
  render() {
    const { tickers } = this.props;
    console.log(tickers);
    return (
      <div>
        {tickers.map((ticker, i) => (
          <TickerCard name={ticker.Name} count={ticker.Observances} key={i} />)
        )}
      </div>
    )
  }
}
