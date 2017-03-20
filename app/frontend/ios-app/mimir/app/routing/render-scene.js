'use strict';
import React, { Component } from 'react';

import * as names from './route-names';
import WatchlistContainer from '../containers/watchlist.container';
import TabMenuContainer from '../containers/tab-menu.container';

const renderScene = (route, navigator) => {
  switch (route.name) {
    case names.MAIN:
      return (<WatchlistContainer navigator={navigator} />);
    case names.COMPANY_PAGE:
      return (<TabMenuContainer navigator={navigator} />);
    default:
      return (<WatchlistContainer navigator={navigator} />);
  }
}

export default render_scene;
