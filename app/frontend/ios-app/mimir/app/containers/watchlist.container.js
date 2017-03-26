'use strict'
import React, { Component } from 'react'
import { Platform } from 'react-native'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import Watchlist from '../components/watchlist'
import Loading from '../components/loading'

import * as userActions from '../ducks/user'
import * as stockActions from '../ducks/stocks'
import * as twitterDataActions from '../ducks/twitter-data'
import { setActiveTicker } from '../ducks/navigation'
import { logonUser } from '../ducks/logon'

import socket from '../methods/server/socket'

import { persistObject } from './../methods/async-storage'
import { arrayEquals } from '../methods/helper-methods'
import { company_page_route } from '../routing/routes'
import { SERVER_URL } from '../credentials/server-info'

class WatchlistContainer extends Component {
  constructor(props) {
    super(props)
    this.socket = socket
  }

  componentWillMount() {
    const { logonUser, reciveTwitterData, updateStockData } = this.props.actions
    logonUser(this.socket)
    this.socket.on("DISPATCH_TWITTER_DATA", payload => {
      if (payload.data) { reciveTwitterData(payload.data) }
    })
    setInterval(() => {
      updateStockData(this.props.state.user.tickers)
    }, 300000) // set this to 30000 (i.e. 30 s. before changing to relese)
  }

  componentWillReceiveProps(nextProps) {
    const { fetchTwitterData, fetchStockData } = this.props.actions
    const { user: nextUser, stocks: nextStocks } = nextProps.state
    const { user } = this.props.state
    const { socket } = this

    if (!user.loaded && nextUser.tickers.length) {
      socket.on("NEW_TWITTER_DATA", () => {
        fetchTwitterData(nextUser, socket)
      })
    } else if (user.loaded && !arrayEquals(nextUser.tickers, user.tickers)) {
      // This happens when the user has user has added or remove a ticker
      const { id, tickers } = nextUser
      persistObject("user", { id, tickers })
      socket.removeListener("NEW_TWITTER_DATA")
      socket.on("NEW_TWITTER_DATA", () => {
        fetchTwitterData(nextUser, socket)
      })
      fetchTwitterData(nextUser, socket)
      fetchStockData(tickers)
    }
  }

  navigateToCompany = ticker => {
    const { navigator, actions } = this.props
    actions.setActiveTicker(ticker)
    navigator.push(company_page_route(ticker))
  }

  removeTicker = ticker => {
    this.props.actions.remove_ticker(ticker)
  }

  render() {
    const { user, stocks } = this.props.state
    if (user.loaded && stocks.loaded) {
      return (
        <Watchlist
          user={user}
          stocks={stocks}
          removeTicker={this.removeTicker}
          navigate={this.navigateToCompany}
        />
      )
    } else {
      return (<Loading />)
    }
  }
}

export default connect(
  (state) => ({
    state: {
      user: state.user,
      stocks: state.stocks,
      navigation: state.navigation
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      ...userActions,
      ...stockActions,
      ...twitterDataActions,
      setActiveTicker,
      logonUser
    }, dispatch)
  })
)(WatchlistContainer)
