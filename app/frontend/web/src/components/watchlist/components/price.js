import React, { Component } from 'react';
import { length, font, color } from '../../../styles/main';
import { round, formatPriceChange, isPositive } from '../../../methods/helper-methods';

const style = {
  container: {
    float: 'right',
    paddingRight: length.medium
  },
  change: {
    fontSize: font.size.medium,
    textAlign: 'right',
    fontWeight: 'bold',
    marginBottom: length.small
  },
  price: {
    fontSize: font.size.normal,
    color: color.grey.dark,
    textAlign: 'right',
    marginTop: '0'
  }
}

export default class Price extends Component {
  getChangeStyle = () => ({
    ...style.change,
    color: (isPositive(this.props.percentChange)) ? color.green : color.red
  })

  render = () => {
    const { percentChange, price, currency } = this.props;
    return (
      <div style={style.container}>
        <p style={this.getChangeStyle()}>{formatPriceChange(percentChange)}</p>
        <p style={style.price}>{round(price)} {currency}</p>
      </div>
    )
  }
}
