'use strict'
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import Watchlist from '../components/watchlist';
import Loading from '../components/loading';

import * as userActions from '../ducks/user';
import * as stockActions from '../ducks/stocks';
import * as twitterDataActions from '../ducks/twitter-data';
import { setActiveTicker } from '../ducks/navigation';
import { logonUser } from '../ducks/logon';

import { remove, USER_ID_KEY } from './../methods/async-storage';
import { arrayEquals } from '../methods/helper-methods';
import { companyPageRoute } from '../routing/main';
import { DEV_MODE } from '../credentials/config';

class WatchlistContainer extends Component {
  componentWillMount() {
    const { logonUser, updateStockData } = this.props.actions;
    //remove(USER_ID_KEY);
    logonUser();
    setInterval(() => {
      updateStockData(this.props.state.user.tickers);
    }, (!DEV_MODE) ? 30000 : 300000); // set this to 30000 (i.e. 30 s. before changing to relese)*/
  }

  componentWillReceiveProps(nextProps) {
    const { fetchTwitterData, fetchStockData } = this.props.actions;
    const { tickers: newTickers } = nextProps.state.user;
    const { user } = this.props.state;
    if (user.loaded && !arrayEquals(newTickers, user.tickers)) {
      // This happens when the user has user has added or remove a ticker
      fetchTwitterData(newTickers);
      fetchStockData(newTickers);
    }
  }

  navigateToCompany = ticker => {
    const { navigator, actions } = this.props;
    actions.setActiveTicker(ticker);
    navigator.push(companyPageRoute(ticker));
  }

  removeTicker = ticker => {
    this.props.actions.removeTicker(ticker);
  }

  userAndDataLoaded = () => {
    const { user, stocks, twitterData } = this.props.state;
    return (user.loaded && stocks.loaded && twitterData.loaded);
  }

  render() {
    const { user, stocks, twitterData } = this.props.state;
    if (this.userAndDataLoaded()) {
      return (
        <Watchlist
          user={user}
          stocks={stocks}
          twitterData={twitterData}
          removeTicker={this.removeTicker}
          navigate={this.navigateToCompany}
        />
      );
    } else {
      return (<Loading />);
    }
  }
}

const mapStateToProps = state => ({
  state: {
    user: state.user,
    stocks: state.stocks,
    twitterData: state.twitterData,
    navigation: state.navigation
  }
})

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    ...userActions,
    ...stockActions,
    ...twitterDataActions,
    setActiveTicker,
    logonUser
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(WatchlistContainer);
