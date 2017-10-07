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
    marginBottom: length.small
  }
}

export default class StockCard extends Component {
  handleClick = () => {
    const { Symbol, navigate } = this.props;
    navigate(Symbol);
  }

  render() {
    const {
      Name: StockName,
      Symbol,
      PercentChange,
      LastTradePriceOnly,
      Currency,
      twitterData
    } = this.props;
    return (
      <div onClick={this.handleClick} className="card" style={style.card}>
        <UrgencyInicator {...twitterData} />
        <Name
          name={StockName}
          ticker={Symbol} />
        <Price
          percentChange={PercentChange}
          price={LastTradePriceOnly}
          currency={Currency} />
      </div>
    )
  }
}
