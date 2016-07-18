'use strict';

import React, { Component } from 'react';
import { Navigator, Text, StyleSheet } from 'react-native';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';

import * as reducers from '../reducers';
import MimirApp from './main.container';

const logger = createLogger()
const createStoreWithMiddleware = applyMiddleware(thunk, logger)(createStore);
const reducer = combineReducers(reducers);
const store = createStoreWithMiddleware(reducer);

export default class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <Navigator
          initialRoute = {{ title: 'mimir', index: 0 }}
          style = {styles.container}
          renderScene = {(route, navigator) => <MimirApp />}
          navigationBar={
             <Navigator.NavigationBar
               routeMapper={{
                 LeftButton: (route, navigator, index, navState) => (<Text>Back</Text>),
                 RightButton: (route, navigator, index, navState) => (<Text>Search</Text>),
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
