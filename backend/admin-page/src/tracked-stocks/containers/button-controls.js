import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { untrackStock, toggleEditMode } from '../../actions/stock-actions';
import ButtonControls from '../components/button-controls';

class ButtonControlsContainer extends Component {
  untrackStock = () => {
    const { actions, state, Ticker } = this.props;
    actions.untrackStock(Ticker, state.token)
  }

  toggleEdit = () => {
    const { actions, Ticker } = this.props
    actions.toggleEditMode(Ticker)
  }

  render() {
    return (
      <ButtonControls
        untrackStock={this.untrackStock}
        toggleEdit={this.toggleEdit}
        editMode={this.props.editMode}
      />
    )
  }
}

export default connect(
  state => ({
    state: {
      token: state.user.token
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      untrackStock,
      toggleEditMode
    }, dispatch)
  })
)(ButtonControlsContainer);
