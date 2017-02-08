import React, { Component } from 'react';
import MainMenu from '../../main-menu/components/main-menu';
import PageTitle from '../../components/util/page-title';
import Loading from '../../components/util/loading';
import Stocklist from './stocklist';
import FilterBarContainer from '../containers/filter-bar';
import { filterStocks } from '../../methods/helper-methods';

export default class TrackedStocks extends Component {
  render() {
    const { loaded, data, filter } = this.props;
    const contents = (loaded)
    ? <Stocklist stocks={filterStocks(data, filter)}/>
    : <Loading/>
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content'>
          <PageTitle title={"Tracked Stocks"} />
          <FilterBarContainer />
          {contents}
        </div>
      </div>
    )
  }
}
