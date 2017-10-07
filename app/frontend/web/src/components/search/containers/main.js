import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { addNewTicker } from '../../../ducks/user';
import { updateAndRunQuery, fetchSearchSugestions } from '../../../ducks/search';
import Search from '../components/main';

class SearchContainer extends Component {
  componentDidMount() {
    const { actions, state } = this.props;
    actions.fetchSearchSugestions(state.user.tickers);
  }

  render() {
    const { search, user } = this.props.state;
    const { addNewTicker, updateAndRunQuery } = this.props.actions
    return (
      <Search
        {...search}
        userId={user.id}
        tickers={user.tickers}
        addNewTicker={addNewTicker}
        updateAndRunQuery={updateAndRunQuery} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    user: state.user,
    search: state.search,
  }
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    addNewTicker,
    updateAndRunQuery,
    fetchSearchSugestions
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToProps(dispatch)
)(SearchContainer)
