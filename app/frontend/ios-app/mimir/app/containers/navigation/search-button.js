'use strict'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { toggleSearchActive } from '../../ducks/search'
import SearchButton from '../../components/navigation/search-button'

class SearchButtonContainer extends Component {
  toggleSearch = () => {
    this.props.actions.toggleSearchActive()
  }

  render() {
    const { props, toggleSearch } = this
    return (
      <SearchButton active={props.state.search.active} action={toggleSearch} />
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
