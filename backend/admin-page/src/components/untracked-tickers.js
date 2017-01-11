import React, { Component } from 'react';
import MainMenu from './main-menu';
import PageTitle from './util/page-title';
import TickerList from './untracked-tickers/ticker-list';

export default class UntrackedTickers extends Component {
  constructor(props) {
    super(props);
    this.state = {
      tickers: null
    }
  }

  componentDidMount() {
    fetch("http://localhost:8080/untracked-tickers", {mode: 'cors'})
    .then(res => res.json())
    .then(resBody => {
      this.setState({
        tickers: resBody
      })
    })
    .catch(err => {console.log(err)});
  }

  render() {
    const { tickers } = this.state;
    const list = (tickers) ? <TickerList tickers={tickers}/> : <p>Loading...</p>
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
