import React, { Component } from 'react';
import { browserHistory } from 'react-router';
import { length, font } from '../../styles/styles';

const styles = {
  listItem: {
    marginTop: length.medium,
    paddingBottom: length.medium,
    fontSize: font.size.medium,
    borderBottom: 'solid',
    borderColor: 'rgba(0,0,0,0.1)',
    borderWidth: length.tiny
  }
}


export default class MenuItem extends Component {
  handleClick = () => {
    const { type, path } = this.props;
    if (type !== 'external') {
      browserHistory.push(this.props.path)
    } else {
      window.location.replace(path)
    }
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
