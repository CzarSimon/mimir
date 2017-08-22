'use strict'
import React, { Component } from 'react';
import { Text, StyleSheet } from 'react-native';
import { Navigator } from 'react-native-deprecated-custom-components';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';

import * as reducers from '../ducks';
import BackButtonContainer from './navigation/back-button';
import SearchButtonContainer from './navigation/search-button';
import TitleContainer from './navigation/title';
import renderScene, { MAIN_ROUTE } from '../routing/main';
import { DEV_MODE } from '../credentials/config';
import SplashScreen from 'react-native-splash-screen';

const logger = createLogger();
const createStoreWithMiddleware = (!DEV_MODE)
  ? applyMiddleware(thunk)(createStore)
  : applyMiddleware(thunk, logger)(createStore)
const reducer = combineReducers(reducers);
const store = createStoreWithMiddleware(reducer);

export default class App extends Component {
  componentDidMount() {
    SplashScreen.hide();
  }

  render() {
    return (
      <Provider store={store}>
        <Navigator
          initialRoute = { MAIN_ROUTE }
          style = { styles.container }
          renderScene = { renderScene }
          navigationBar = {
            <Navigator.NavigationBar
              routeMapper={{
                LeftButton: (route, navigator, index) => (<BackButtonContainer navigator={navigator} index={index}/>),
                RightButton: (route, navigator, index) => (<SearchButtonContainer route={route} navigator={navigator} index={index}/>),
                Title: (route) => (<TitleContainer route={route} />)
              }}
              style={{borderBottomWidth: 1, borderColor: 'lightgrey'}}
            />
          }
        />
      </Provider>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'stretch',
    backgroundColor: "#F8F8F8"
  }
});
