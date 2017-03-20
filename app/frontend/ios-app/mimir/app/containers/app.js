'use strict';

import React, { Component } from 'react'
import { Navigator, Text, StyleSheet } from 'react-native'
import { createStore, applyMiddleware, combineReducers } from 'redux'
import { Provider } from 'react-redux'
import thunk from 'redux-thunk'
import createLogger from 'redux-logger'

import * as reducers from '../ducks'
import BackButton from './navigation/back-button'
import SearchButtonContainer from './navigation/search-button'
import TitleContainer from './navigation/title'
import renderScene from '../routing/render-scene'
import { MAIN_ROUTE } from '../routing/routes'
import { DEV_MODE } from '../credentials/config'

const logger = createLogger();
const createStoreWithMiddleware = (!DEV_MODE)
  ? applyMiddleware(thunk)(createStore)
  : applyMiddleware(thunk, logger)(createStore)
const reducer = combineReducers(reducers);
const store = createStoreWithMiddleware(reducer);

export default class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <Navigator
          initialRoute = {MAIN_ROUTE}
          style = {styles.container}
          renderScene = {renderScene}
          navigationBar = {
            <Navigator.NavigationBar
              routeMapper={{
                LeftButton: (route, navigator, index, navState) => (<BackButton index={index} nav={navigator}/>),
                RightButton: () => (<SearchButtonContainer />),
                Title: (route) => (<TitleContainer title={route.title} />)
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
