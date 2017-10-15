import React, { Component } from 'react';
import { color, length } from '../../../styles/main';

const style = {
  container: {
    textAlign: 'center',
    alignSelf: 'center',
    border: 'solid',
    borderWidth: '1px',
    borderRadius: '5px',
    padding: length.small,
    marginBottom: length.small,
    color: color.grey.dark
  }
}

export default class UnfollowButton extends Component {
  render() {
    const { unfollowTicker } = this.props;
    return (
      <div
        className='button'
        style={style.container}
        onClick={unfollowTicker}>
        Unfollow
      </div>
    )
  }
}
