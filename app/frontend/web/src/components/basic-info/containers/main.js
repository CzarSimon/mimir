import React, { Component } from 'react'
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchNewsItems } from '../../../ducks/news';
import { setActiveTicker } from '../../../ducks/navigation';
import BasicInfo from '../components/main';

class BasicInfoContainer extends Component {
  render() {
    const { navigation, stocks } = this.props.state;
    if (navigation.activeTicker && stocks.loaded) {
      const {
        name,
        price,
        currency,
        priceChange
      } = stocks.data[navigation.activeTicker];
      return (
        <BasicInfo
          name={name}
          price={price}
          currency={currency}
          priceChange={priceChange}/>
      );
    } else {
      return <div />
    }
  }
}

const mapStateToProps = state => ({
  state: {
    news: state.news,
    stocks: state.stocks,
    navigation: state.navigation
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    fetchNewsItems,
    setActiveTicker
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(BasicInfoContainer);
