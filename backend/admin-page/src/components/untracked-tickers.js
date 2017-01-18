import React, { Component } from 'react';
import MainMenu from './main-menu';
import PageTitle from './util/page-title';
import TickerList from './untracked-tickers/ticker-list';

export default class UntrackedTickers extends Component {
  render() {
    const { tickers, loaded } = this.props;
    console.log('tickers', tickers);
    const list = (loaded) ? <TickerList tickers={tickers}/> : <p>Loading...</p>
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content'>
          <PageTitle title={"Untracked tickers"} />
          {list}
        </div>
      </div>
    )
  }
}
