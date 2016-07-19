'use strict';
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { add_ticker } from '../actions/user.actions';
import { fetch_company_desc } from '../actions/descriptions.actions';
import Overview from '../components/overview';
import { arr_get_value_by_key, format_name } from '../methods/helper-methods';

class OverviewContainer extends Component {
  componentWillMount() {
    const { company, actions, state } = this.props;
    const { Name, Symbol } = company;
    if (!state.descriptions[company.Symbol]) {
      actions.fetch_company_desc(format_name(Name), Symbol);
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
      fetch_company_desc,
      add_ticker
    }, dispatch)
  })
)(OverviewContainer);
