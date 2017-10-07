import React, { Component } from 'react';
import { length } from '../../../styles/main';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { addNewTicker } from '../../../ducks/user';
import { updateAndRunQuery, fetchSearchSugestions } from '../../../ducks/search';
//import SearchHistory from '../components/search-history';
//import SearchSugestions from '../components/search-sugestions';
//import SearchResults from '../components/search-results';
import Search from '../components/main';

class SearchContainer extends Component {
  componentDidMount() {
    const { actions, state } = this.props;
    actions.fetchSearchSugestions(state.user.tickers);
  }

  goToStock = (ticker, added = true) => {
    const { navigator } = this.props;
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
