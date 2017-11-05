import React, { Component } from 'react';
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
  renderItem = (ticker, key) => {
    const { twitterData, navigate, stocks } = this.props;
    const stock = stocks.data[ticker]
    return (
      <StockCard
        key={key}
        {...stock}
        className='list-group-item'
        twitterData={twitterData[ticker]}
        navigate={navigate} />
    )
  }

  render() {
    const { logout, user } = this.props;
    return (
      <div>
        <p style={style.header}>Watchlist</p>
        {user.tickers.map(this.renderItem)}
        <LogoutButton logout={logout}/>
      </div>
    )
  }
}
