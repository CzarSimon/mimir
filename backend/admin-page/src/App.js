import React, { Component } from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { syncHistoryWithStore, routerReducer } from 'react-router-redux';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';
// have to install redux-thunk and react-redux!!!


import * as reducers from './reducers';
import TrackedStocks from './components/tracked-stocks';
import UntrackedTickers from './components/untracked-tickers';
import UntrackedInfo from './components/untracked-info';
//import './App.css';


const logger = createLogger();
const createStoreWithMiddleware = applyMiddleware(logger, thunk)(createStore);
const reducer = combineReducers({
  ...reducers,
  routing: routerReducer
});
const store = createStoreWithMiddleware(reducer);
const history = syncHistoryWithStore(browserHistory, store);


export default class App extends Component {
  render() {

    return (
      <Provider store={store}>
        <Router history={history}>
          <Route path="/" component={TrackedStocks} />
          <Route path="/tracked-stocks" component={TrackedStocks} />
          <Route path="/untracked-tickers" component={UntrackedTickers} />
          <Route path="/ticker/:tickerName" component={UntrackedInfo} />
        </Router>
      </Provider>
    );
  }
}
