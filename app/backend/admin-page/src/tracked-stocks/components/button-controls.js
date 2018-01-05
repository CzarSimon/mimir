import React, { Component } from 'react';
import Button from '../../components/util/button';
import { length, color } from '../../styles/styles';

const styles = {
  untrack: {
    backgroundColor: color.red,
    clear: 'both',
    float: 'right'
  },
  edit: {
    paddingLeft: length.large,
    paddingRight: length.large
  },
  save: {
    paddingLeft: length.large,
    paddingRight: length.large,
    backgroundColor: color.green
  }
}

export default class ButtonControls extends Component {
  editClick = () => {
    this.props.toggleEdit()
  }

  untrackClick = () => {
    this.props.untrackStock()
  }

  render() {
    const editActionButton = (!this.props.editMode)
    ? <Button text={'Edit'} customStyles={styles.edit} handleClick={this.editClick}/>
    : <Button text={'Cancel'} customStyles={styles.save} handleClick={this.editClick}/>
    return (
      <div>
        <Button
          text={'Untrack stock'}
          customStyles={styles.untrack}
          handleClick={this.untrackClick}
        />
        {editActionButton}
      </div>
    )
  }
}
