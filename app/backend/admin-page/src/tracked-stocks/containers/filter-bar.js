import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { updateFilter } from '../../actions/stock-actions';
import FilterBar from '../components/filter-bar';


class FilterBarContainer extends Component {
  handleFilterInput = filterTerm => {
    this.props.actions.updateFilter(filterTerm)
  }

  render() {
    return (
      <FilterBar
        filterTerm={this.props.state.filterTerm}
        handleFilterInput={this.handleFilterInput}
      />
    )
  }
}

export default connect(
  state => ({
    state: {
      filterTerm: state.stocks.filterTerm
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      updateFilter
    }, dispatch)
  })
)(FilterBarContainer);
