import React, { Component } from 'react';
import TickerCard from './ticker-card';

export default class TickerList extends Component {
  render() {
    const untrackedTickers = [
      {"name": "NVDA", "observances": 321},
      {"name": "PG", "observances": 78},
      {"name": "GM", "observances": 210},
      {"name": "F", "observances": 33},
      {"name": "CRM", "observances": 421},
    ]

    return (
      <div>
        {untrackedTickers.map((ticker, i) => (
          <TickerCard name={ticker.name} key={i} />)
        )}
      </div>
    )
  }
}
