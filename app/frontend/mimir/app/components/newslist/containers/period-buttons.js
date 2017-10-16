import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchNewsItems, switchPeriod } from '../../../ducks/news';
import PeriodButtons from '../components/period-buttons';

class PeriodButtonsContainer extends Component {
  handleClick = period => {
    const { switchPeriod, fetchNewsItems } = this.props.actions;
    const { ticker } = this.props.state;
    switchPeriod(period);
    fetchNewsItems(ticker, period);
  }
  
  render() {
    const { period } = this.props.state.news;
    return <PeriodButtons period={period} select={this.handleClick} />
  }
}

const mapStateToProps = state => ({
  state: {
    news: state.news,
    ticker: state.navigation.activeTicker
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    fetchNewsItems,
    switchPeriod
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(PeriodButtonsContainer);
