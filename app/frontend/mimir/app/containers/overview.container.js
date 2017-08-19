'use strict'

import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { addTicker } from '../ducks/user';
import { fetchStockData } from '../ducks/stocks'
import { fetchCompanyDesc } from '../ducks/descriptions';
import Overview from '../components/overview';
import { formatName } from '../methods/helper-methods';

class OverviewContainer extends Component {
  componentDidMount() {
    const { company, actions, state } = this.props;
    const { Name, Symbol } = company;
    if (!state.descriptions[company.Symbol]) {
      actions.fetchCompanyDesc(formatName(Name), Symbol);
    }
    actions.fetchStockData([ Symbol ]);
  }

  render() {
    const { company, twitterData, state }  = this.props;
    const description = state.descriptions[company.Symbol];
    return (
      <Overview
        twitterData={twitterData}
        description={description}
        company={company}
      />
    );
  }
}

const mapStateToProps = state => ({
  state: {
    user: state.user,
    descriptions: state.descriptions
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    fetchCompanyDesc,
    fetchStockData,
    addTicker
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(OverviewContainer);
