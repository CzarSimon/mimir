import React, { Component } from 'react'
import UnfollowButtonContainer from '../containers/unfollow-button';
import { round, formatPriceChange, isPositive } from '../../../methods/helper-methods';
import { length, color, font } from '../../../styles/main';

const style = {
  container: {
    display: 'flex',
    border: 'solid',
    borderWidth: 0,
    borderBottomWidth: '1px',
    borderColor: color.blue
  },
  info: {
    flex: 5,
    paddingLeft: length.small
  },
  name: {
    color: color.blue,
    fontSize: font.size.small,
    fontWeight: 'bold',
    margin: 0
  },
  price: {
    display: 'flex',
    marginTop: length.mini,
    marginBottom: length.mini,
    color: color.black
  },
  priceText: {
    margin: 0
  },
  changeText: {
    margin: 0,
    marginLeft: length.small
  }
}

export default class BasicInfo extends Component {
  render() {
    const { name, price, currency, priceChange } = this.props;
    const changeColor = (isPositive(priceChange)) ? color.green : color.red;
    return (
      <div style={style.container}>
        <div></div>
        <div style={style.info}>
          <p style={style.name}>{name}</p>
          <div style={style.price}>
            <p style={style.priceText}>{round(price)} {currency}</p>
            <p style={{...style.changeText, color: changeColor}}>
              {formatPriceChange(priceChange)}
            </p>
          </div>
        </div>
        <UnfollowButtonContainer />
      </div>
    );
  }
}
