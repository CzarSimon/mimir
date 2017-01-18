import React, { Component } from 'react';
import MainMenu from './main-menu';
import InfoName from './util/info-name';
import { length } from '../styles/styles';

const styles = {
  card: {
    paddingTop: length.medium
  },
  text: {
    paddingBottom: length.mini
  }
}

export default class UntrackedInfo extends Component {
  render() {
    const description = "This is the description of the company that has this ticker"
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content untracked-info' style={styles.card}>
          <InfoName name={"Stock name"} />
          <p style={styles.text}>{this.props.params.tickerName}</p>
          <h3 style={styles.text}>Description</h3>
          <p style={styles.text}>{description}</p>
          <button>Track ticker</button>
        </div>
      </div>
    )
  }
}
