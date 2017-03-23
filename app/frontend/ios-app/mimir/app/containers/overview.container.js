'use strict'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { addTicker } from '../ducks/user'
import { fetchCompanyDesc } from '../ducks/descriptions'
import Overview from '../components/overview'
import { formatName } from '../methods/helper-methods'

class OverviewContainer extends Component {
  componentWillMount() {
    const { company, actions, state } = this.props;
    const { Name, Symbol } = company;
    if (!state.descriptions[company.Symbol]) {
      actions.fetchCompanyDesc(formatName(Name), Symbol);
    }
  }
  render() {
    const { company, twitter_data, state }  = this.props;
    const description = state.descriptions[company.Symbol]
    return (
      <Overview
        twitter_data = {twitter_data}
        description = {description}
        company = {company}
      />
    );
  }
}

export default connect(
  (state) => ({
    state: {
      user: state.user,
      descriptions: state.descriptions
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      fetchCompanyDesc,
      addTicker
    }, dispatch)
  })
)(OverviewContainer);
