import React, { Component } from 'react';
import { Router, Route, browserHistory } from 'react-router';
import TrackedStocks from './components/tracked-stocks';
import UntrackedTickers from './components/untracked-tickers';
import UntrackedInfo from './components/untracked-info';
//import './App.css';

export default class App extends Component {
  render() {
    return (
      <Router history={browserHistory}>
        <Route path="/" component={TrackedStocks} />
        <Route path="/tracked-stocks" component={TrackedStocks} />
        <Route path="/untracked-tickers" component={UntrackedTickers} />
        <Route path="/ticker/:tickerName" component={UntrackedInfo} />
      </Router>
    );
  }
}
