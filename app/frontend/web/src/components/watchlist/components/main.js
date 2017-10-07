import React, { Component } from 'react';
import _ from 'lodash';
import StockCard from './stock-card';
import LogoutButton from '../../logout/components/main';
import { font, color, length } from '../../../styles/main';

const style = {
  header: {
    fontSize: font.size.large,
    color: color.blue,
    paddingLeft: length.small,
    marginTop: length.small,
    marginBottom: length.small
  }
}

export default class Watchlist extends Component {
  renderItem = (stock, key) => {
    const { twitterData, navigate } = this.props;
    return (
      <StockCard
        key={key}
        {...stock}
        className='list-group-item'
        twitterData={twitterData[stock.Symbol]}
        navigate={navigate} />
    )
  }

  render() {
    const { stocks, logout } = this.props;
    return (
      <div>
        <p style={style.header}>Watchlist</p>
        {_.values(stocks.data).map(this.renderItem)}
        <LogoutButton logout={logout}/>
      </div>
    )
  }
}
