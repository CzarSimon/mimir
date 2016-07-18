'use strict';

import React, { Component } from 'react';
import { Navigator, Text, StyleSheet } from 'react-native';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';

import * as reducers from '../reducers';
import BackButton from './navigation/back-button';
import SearchButton from './navigation/search-button';
import render_scene from '../routing/render-scene';
import { MAIN_ROUTE } from '../routing/routes';

const logger = createLogger();
const createStoreWithMiddleware = applyMiddleware(thunk, logger)(createStore);
const reducer = combineReducers(reducers);
const store = createStoreWithMiddleware(reducer);

export default class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <Navigator
          initialRoute = {MAIN_ROUTE}
          style = {styles.container}
          renderScene = {render_scene}
          navigationBar={
            <Navigator.NavigationBar
              routeMapper={{
                LeftButton: (route, navigator, index, navState) => (<BackButton index={index} nav={navigator}/>),
                RightButton: () => (<SearchButton active={false} />),
                Title: (route, navigator, index, navState) => (<Text>{route.title}</Text>)
              }}
              style={{backgroundColor: 'gray'}}
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
    alignItems: 'stretch'
  }
});
