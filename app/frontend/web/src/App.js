import React, { Component } from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { syncHistoryWithStore, routerReducer } from 'react-router-redux';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import logger from 'redux-logger';

import * as reducers from './ducks';
import { DEV_MODE } from './credentials/config';
import { portraitMode } from './methods/helper-methods';
import Auth from './methods/auth-service';
import HeaderContainer from './components/header/containers/main';
import WatchlistContainer from './components/watchlist/containers/main';
import NewslistContainer from './components/newslist/containers/main';
import SearchContainer from './components/search/containers/main';
import CallbackContainer from './components/login/containers/main';
import Login from './components/login/components/main';
import { updateStockData } from './ducks/stocks';
import { fetchTwitterData } from './ducks/twitter-data';
import { logonUser } from './ducks/logon';

const createStoreWithMiddleware = (!DEV_MODE && false)
  ? applyMiddleware(thunk)(createStore)
  : applyMiddleware(thunk, logger)(createStore)
const reducer = combineReducers({
  ...reducers,
  routing: routerReducer
});
const store = createStoreWithMiddleware(reducer);
const history = syncHistoryWithStore(browserHistory, store);

const style = {
  display: 'table',
  margin: '60px auto',
  width: (portraitMode()) ? '94%' : '50%'
}

const updateFrequency = {
  DEV: 300000,
  PROD: 30000
}

class App extends Component {
  constructor(props) {
    super(props);
    this.auth = new Auth();
    this.state = {
      width: (portraitMode()) ? '94%' : '50%'
    }
  }

  updateDimensions = () => {
    this.setState({
      width: (portraitMode()) ? '94%' : '50%'
    });
  }

  componentDidMount() {
    if (this.auth.isAuthenticated()) {
      store.dispatch(logonUser());
    }
    setInterval(() => {
      const { tickers } = store.getState().user;
      store.dispatch(updateStockData(tickers));
      store.dispatch(fetchTwitterData(tickers));
    }, (!DEV_MODE) ? updateFrequency.PROD : updateFrequency.DEV);
    window.addEventListener("resize", this.updateDimensions);
  }

  componentWillUnmount() {
    window.removeEventListener("resize", this.updateDimensions);
  }

  requireAuth = () => {
    if (!this.auth.isAuthenticated()) {
      browserHistory.push('/login');
    }
  }

  renderCallbackHandler = () => {
    return <CallbackContainer auth={this.auth} />
  }

  renderLogin = () => (
    <Login auth={this.auth} />
  )

  renderWatchlist = () => (
    <WatchlistContainer logout={this.auth.logout} />
  )

  render() {
    const { width } = this.state;
    return (
      <Provider store={store}>
        <div className="App">
          <HeaderContainer />
          <div style={{...style, width}}>
            <Router history={history}>
              <Route path="/login" component={this.renderLogin} />
              <Route
                path="/callback"
                component={this.renderCallbackHandler} />
              <Route
                path="/"
                onEnter={this.requireAuth}
                component={this.renderWatchlist} />
              <Route
                path="/news/:ticker"
                onEnter={this.requireAuth}
                component={NewslistContainer} />
              <Route
                path="/search"
                onEnter={this.requireAuth}
                component={SearchContainer} />
            </Router>
          </div>
        </div>
      </Provider>
    );
  }
}

export default App;
