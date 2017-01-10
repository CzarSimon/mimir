import React, { Component } from 'react';
import MainMenu from './main-menu';
import PageTitle from './util/page-title';
import TickerList from './untracked-tickers/ticker-list';

export default class UntrackedTickers extends Component {
  render() {
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content'>
          <PageTitle title={"Untracked tickers"} />
          <TickerList />
        </div>
      </div>
    )
  }
}
