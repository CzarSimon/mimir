'use strict';
import React, { Component } from 'react';

import * as names from './route-names';
import MimirApp from '../containers/main.container';
import TabMenuContainer from '../containers/tab-menu.container';

const render_scene = (route, navigator) => {
  switch (route.name) {
    case names.MAIN:
      return (<MimirApp navigator={navigator} />);
    case names.COMPANY_PAGE:
      return (<TabMenuContainer navigator={navigator} />);
    default:
      return (<MimirApp navigator={navigator} />);
  }
}

export default render_scene;
