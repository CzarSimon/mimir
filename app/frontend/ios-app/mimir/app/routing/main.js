'use strict'
import React, { Component } from 'react'
import WatchlistContainer from '../containers/watchlist.container'
import TabMenuContainer from '../containers/tab-menu.container'
import SearchContainer from '../search/containers/main'
import { last } from 'lodash'

/* --- Route constants --- */
export const MAIN = 'MAIN'
export const COMPANY_PAGE = 'COMPANY_PAGE'
export const SEARCH_PAGE = 'SEARCH_PAGE'


/* --- Route switcher --- */
const renderScene = (route, navigator) => {
  switch (route.name) {
    case MAIN:
      return <WatchlistContainer navigator={navigator} />
    case COMPANY_PAGE:
      return <TabMenuContainer navigator={navigator} />
    case SEARCH_PAGE:
      return <SearchContainer navigator={navigator} />
    default:
      return <WatchlistContainer navigator={navigator} />
  }
}
export default renderScene


/* --- Routes --- */
export const MAIN_ROUTE = {
  name: MAIN,
  title: 'mimir',
  index: 0
}


export const getSearchRoute = currentIndex => (
  {
    name: SEARCH_PAGE,
    title: 'search',
    index: currentIndex++
  }
)


export const companyPageRoute = (title, newIndex = 1) => (
  {
    name: COMPANY_PAGE,
    index: newIndex,
    title
  }
)

/* --- Helper methods --- */
export const currentRoute = navigator => last(navigator.getCurrentRoutes())

export const getRouteIndex = navigator => currentRoute.index
