'use strict'
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import BackButton from '../../components/navigation/back-button';
import { cancelSearch } from '../../ducks/search';

class BackButtonContainer extends Component {
  handleClick = () => {
    const { navigator, actions } = this.props;
    actions.cancelSearch();
    navigator.pop();
  }

  render() {
    const { index } = this.props
    return <BackButton index={index} handleClick={this.handleClick} />
  }
}

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    cancelSearch
  }, dispatch)
});

export default connect(
  state => ({}),
  dispatch => mapDispatchToActions(dispatch)
)(BackButtonContainer)
