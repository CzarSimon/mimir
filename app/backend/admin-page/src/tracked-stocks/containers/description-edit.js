import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { updateStockInfo, toggleEditMode } from '../../actions/stock-actions'
import DescriptionEdit from '../components/description-edit'


class DescriptionEditContainer extends Component {
  updateInfo = description => {
    const { actions, state, Ticker } = this.props
    actions.updateStockInfo(Ticker, description, state.token)
  }

  toggleEdit = () => {
    const { actions, Ticker } = this.props
    actions.toggleEditMode(Ticker)
  }

  render() {
    const { currentDescription } = this.props
    return (
      <DescriptionEdit
        currentDescription={currentDescription}
        updateInfo={this.updateInfo}
        toggleEdit={this.toggleEdit}
      />
    )
  }
}


const mapStateToProps = state => ({
  state: {
    token: state.user.token
  }
})


const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    updateStockInfo,
    toggleEditMode
  }, dispatch)
})


export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToProps(dispatch)
)(DescriptionEditContainer)
