import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { selectMenuItem } from '../../actions/menu-actions';
import MenuItem from '../components/menu-item';

class MenuItemContainer extends Component {
  handleClick = () => {
    const { actions, state, type, path, name: menuItem } = this.props;
    if (state.selected !== menuItem) {
      actions.selectMenuItem(menuItem)
    }
    if (type !== 'external') {
      browserHistory.push(path)
    } else {
      window.location.replace(path)
    }
  }

  render() {
    const { idName, name, state } = this.props
    const selected = (name === state.selected)
    return (
      <MenuItem
        selected={selected}
        idName={idName}
        name={name}
        handleClick={this.handleClick}
      />
    )
  }
}

export default connect(
  state => ({
    state: {
      selected: state.menu.selected
    }
  }),
  dispatch => ({
    actions: bindActionCreators({
      selectMenuItem
    }, dispatch)
  })
)(MenuItemContainer);
