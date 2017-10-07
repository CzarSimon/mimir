import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { updateAndRunQuery } from '../../../ducks/search';

import SearchBar from '../components/search-bar';

class SearchBarContainer extends Component {
  search = query => {
    this.props.actions.updateAndRunQuery(query);
  }

  render() {
    const { search } = this.props.state;
    return (
      <SearchBar {...search} search={this.search} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    search: state.search
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    updateAndRunQuery
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(SearchBarContainer)
