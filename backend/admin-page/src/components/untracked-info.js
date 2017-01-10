import React, { Component } from 'react';
import MainMenu from './main-menu';
import InfoName from './util/info-name';

export default class UntrackedInfo extends Component {
  render() {
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content untracked-info'>
          <InfoName name={"Stock name"} />
          <p>{this.props.params.tickerName}</p>
          <h4>Description</h4>
          <p>This is the description of the company that has this ticker</p>
          <button>Track ticker</button>
        </div>
      </div>
    )
  }
}
