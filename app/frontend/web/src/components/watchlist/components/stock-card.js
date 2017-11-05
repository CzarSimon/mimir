import React, { Component } from 'react';
import Name from './name';
import Price from './price';
import UrgencyInicator from './urgency-indicator';
import { length } from '../../../styles/main';

const style = {
  card: {
    width: '100%',
    clear: 'both',
    overflow: 'auto',
    marginBottom: length.medium
  }
}

export default class StockCard extends Component {
  handleClick = () => {
    const { ticker, navigate } = this.props;
    navigate(ticker);
  }

  render() {
    const {
      name,
      ticker,
      priceChange,
      price,
      currency,
      twitterData
    } = this.props;
    return (
      <div onClick={this.handleClick} className="card" style={style.card}>
        <UrgencyInicator {...twitterData} />
        <Name
          name={name}
          ticker={ticker} />
        <Price
          percentChange={priceChange}
          price={price}
          currency={currency} />
      </div>
    )
  }
}
