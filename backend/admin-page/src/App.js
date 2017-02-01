import React, { Component } from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { syncHistoryWithStore, routerReducer } from 'react-router-redux';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';

import * as reducers from './reducers';
import TrackedStocks from './components/tracked-stocks';
import UntrackedTickersContainer from './containers/untracked-tickers';
import UntrackedInfoContainer from './containers/untracked-info';
import LoginContainer from './login/containers/login';


const logger = createLogger();
const createStoreWithMiddleware = applyMiddleware(logger, thunk)(createStore);
const reducer = combineReducers({
  ...reducers,
  routing: routerReducer
});
const store = createStoreWithMiddleware(reducer);
const history = syncHistoryWithStore(browserHistory, store);


export default class App extends Component {
  requireAuth = () => {
    const { user } = store.getState()
    if (!user.token) {
      browserHistory.push("/login")
    }
  }

  render() {
    return (
      <Provider store={store}>
        <Router history={history}>
          <Route path="/" onEnter={this.requireAuth} component={TrackedStocks} />
          <Route path="/login" component={LoginContainer} />
          <Route path="/tracked-stocks" onEnter={this.requireAuth} component={TrackedStocks} />
          <Route path="/untracked-tickers" onEnter={this.requireAuth} component={UntrackedTickersContainer} />
          <Route path="/ticker/:tickerName" onEnter={this.requireAuth} component={UntrackedInfoContainer} />
        </Router>
      </Provider>
    );
  }
}
