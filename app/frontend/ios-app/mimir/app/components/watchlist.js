'use strict'

import React, { Component } from 'react'
import { View, Text, StyleSheet, ListView } from 'react-native'
import { length, font, color } from '../styles/styles'
import { values } from 'lodash'

import HeaderContainer from '../containers/watchlist/header.container'
import StockCard from './watchlist/stock-card'
import SearchResultContainer from '../containers/search-result.container'

export default class Watchlist extends Component {
  createUserStockList = (stockData, tickerOrder = []) => {
    return values(stockData)
  }

  render() {
    const { user, stocks, navigate, removeTicker } = this.props
    const twitterData = (user.twitterData.loaded) ? user.twitterData.data : {}
    const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2})
    const userTicketList = ds.cloneWithRows(stocks.data)
    return (
      <View style = {styles.container}>
        <View style={styles.search_result}>
          <SearchResultContainer />
        </View>
        <ListView
          dataSource = {userTicketList}
          renderHeader = {() => <HeaderContainer />}
          style={styles.list}
          renderRow = {(stockData) => {
            if (user.tickers.includes(stockData.Symbol)) {
              return (
                <StockCard
                  {...stockData}
                  twitterData={twitterData[stockData.Symbol]}
                  navigate={navigate}
                  modifiable={user.modifiable}
                  removeTicker={removeTicker}
                  />
              )
            } else {
              return (<View />)
            }
          }}
        />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center',
    marginTop: length.navbar
  },
  list: {
    flex: 1
  },
  header: {
    alignSelf: 'flex-start',
    fontSize: font.h3,
    paddingLeft: length.small,
  },
  search_result: {
    marginLeft: length.medium
  }
})
