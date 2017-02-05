import React, { Component } from 'react';
import MainMenu from '../main-menu/components/main-menu';
import PageTitle from './util/page-title';

export default class TrackedStocks extends Component {
  render() {
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content'>
          <PageTitle title={"Tracked Stocks"} />
          <p>These are the tracked stocks</p>
        </div>
      </div>
    )
  }
}
