import React, { Component } from 'react';
import MainMenu from '../../main-menu/components/main-menu';
import PageTitle from '../../components/util/page-title';
import Loading from '../../components/util/loading';
import Stocklist from './stocklist';
import { values } from 'lodash'

export default class TrackedStocks extends Component {
  render() {
    const { loaded, data } = this.props;
    const contents = (loaded) ? <Stocklist stocks={values(data)}/> : <Loading/>
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content'>
          <PageTitle title={"Tracked Stocks"} />
          {contents}
        </div>
      </div>
    )
  }
}
