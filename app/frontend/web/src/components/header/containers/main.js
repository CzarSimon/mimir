import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import Header from '../components/main';

class HeaderContainer extends Component {
  render() {
    const { pathname, activeTicker } = this.props.state;
    return <Header path={pathname} ticker={activeTicker} />
  }
}

const mapStateToProps = state => ({
  state: {
    pathname: state.routing.locationBeforeTransitions.pathname,
    activeTicker: state.navigation.activeTicker
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({}, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(HeaderContainer);
