import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchNewsItems } from '../../../ducks/news';
import Newslist from '../components/main';
import Loading from '../../helpers/loading';

class NewslistContainer extends Component {
  componentDidMount() {
    const { actions, state } = this.props;
    const { activeTicker } = state.navigation;
    actions.fetchNewsItems(activeTicker, '3M');
  }

  render() {
    const { navigation, news }  = this.props.state;
    const companyNews = news[navigation.activeTicker];
    return (companyNews) ?
      <Newslist news={companyNews} /> :
      <Loading />
  }
}

const mapStateToProps = state => ({
  state: {
    news: state.news,
    navigation: state.navigation
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    fetchNewsItems
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(NewslistContainer);
