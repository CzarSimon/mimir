import React, { Component } from 'react'
import { length, font, color } from '../../styles/styles';
import ButtonGroup from './button-group';

const styles = {
  card: {
    marginBottom: length.medium,
    padding: length.medium,
    backgroundColor: color.white,
    marginRight: '4vw'
  },
  heading: {
    fontSize: font.size.medium,
    marginBottom: length.mini,
    color: color.blue
  },
  text: {
    fontSize: font.size.medium,
    marginBottom: length.medium
  },
  count: {
    fontSize: font.size.medium,
    marginRight: length.mini,
    float: 'right'
  }
}

export default class CandidateCard extends Component {
  render() {
    const { candidate, labelTweet, skip, count } = this.props;
    return (
      <div className='card' style={styles.card}>
        <p style={styles.count}>{count}</p>
        <p style={styles.heading}>Candidate</p>
        <p style={styles.text}>{candidate}</p>
        <ButtonGroup candidate={candidate} labelTweet={labelTweet} skip={skip} />
      </div>
    )
  }
}
