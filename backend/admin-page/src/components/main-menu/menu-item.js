import React, { Component } from 'react';
import { browserHistory } from 'react-router';

const styles = {
  listItem: {
    marginBottom: '15px'
  }
}


export default class MenuItem extends Component {
  handleClick = () => {
    browserHistory.push(this.props.path);
  }

  render() {
    const { idName, name } = this.props;
    return (
      <li style={styles.listItem}>
        <a id={idName} onClick={this.handleClick}>{name}</a>
      </li>
    )
  }
}
