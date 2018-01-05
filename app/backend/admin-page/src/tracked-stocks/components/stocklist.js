import React, { Component } from 'react';
import StockCard from './stock-card';

export default class Stocklist extends Component {
  render() {
    const { stocks } = this.props;
    return (
      <div>
        {stocks.map((stock, i) => <StockCard {...stock} key={i} />)}
      </div>
    )
  }
}
