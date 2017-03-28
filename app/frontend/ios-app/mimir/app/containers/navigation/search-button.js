'use strict'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { toggleSearchActive } from '../../ducks/search'
import SearchButton from '../../components/navigation/search-button'
import { getSearchRoute } from '../../routing/main'

class SearchButtonContainer extends Component {
  toggleSearch = () => {
    const { navigator, actions, index } = this.props
    actions.toggleSearchActive()
    navigator.push(getSearchRoute(index))
  }

  render() {
    const { props, toggleSearch } = this
    return (
      <SearchButton active={props.state.search.active} action={this.toggleSearch} />
    )
  }
}
export default connect(
  (state) => ({
    state: {
      search: state.search
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      toggleSearchActive
    }, dispatch)
  })
)(SearchButtonContainer)
