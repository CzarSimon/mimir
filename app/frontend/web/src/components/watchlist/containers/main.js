import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { updateStockData } from '../../../ducks/stocks';
import { fetchTwitterData } from '../../../ducks/twitter-data';
import { setActiveTicker } from '../../../ducks/navigation';
import { logonUser } from '../../../ducks/logon';

import Watchlist from '../components/main';
import Loading from '../../helpers/loading';

class WatchlistContainer extends Component {
  userAndDataLoaded = () => {
    const { user, stocks, twitterData } = this.props.state;
    return (user.loaded && stocks.loaded && twitterData.loaded);
  }

  selectTicker = ticker => {
    this.props.actions.setActiveTicker(ticker);
    browserHistory.push(`/news/${ticker}`);
  }

  render() {
    const { user, stocks, twitterData } = this.props.state;
    return (this.userAndDataLoaded()) ?
      <Watchlist
        user={user}
        stocks={stocks}
        twitterData={twitterData}
        navigate={this.selectTicker} /> :
      <Loading />
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
    updateStockData,
    fetchTwitterData,
    setActiveTicker,
    logonUser
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(WatchlistContainer);
