import React, { Component } from 'react';
import MainMenu from './main-menu';
import InfoName from './util/info-name';
import { length, color } from '../styles/styles';

const styles = {
  content: {
    paddingTop: length.large
  },
  card: {
    padding: length.large,
    backgroundColor: color.white,
    marginRight: '4vw'
  },
  text: {
    paddingBottom: length.mini
  }
}

export default class UntrackedInfo extends Component {
  render() {
    const { companyName, tickerName, description } = this.props;
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content' style={styles.content}>
          <div className='card' style={styles.card}>
            <InfoName name={companyName} />
            <p style={styles.text}>{tickerName}</p>
            <h3 style={styles.text}>Description</h3>
            <p style={styles.text}>{description}</p>
            <button>Track ticker</button>
          </div>
        </div>
      </div>
    )
  }
}
