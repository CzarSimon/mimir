import React, { Component } from 'react';
import { browserHistory } from 'react-router';
import { color, length } from '../../styles/styles'

const styles = {
  card: {
    backgroundColor: color.white,
    margin: 0,
    marginRight: length.large,
    marginBottom: length.medium,
    padding: length.mini,
    paddingTop: length.small,
    paddingBottom: length.small,
    width: '15%',
    float: 'left'
  },
  text: {
    textAlign: 'center'
  }
}

export default class TickerCard extends Component {
  handleClick = () => {
    browserHistory.push(`/ticker/${this.props.name}`);
  }

  render() {
    return (
      <div className='card' style={styles.card} onClick={this.handleClick}>
        <p style={styles.text}>{this.props.name}</p>
      </div>
    )
  }
}
