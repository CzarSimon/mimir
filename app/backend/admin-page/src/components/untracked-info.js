import React, { Component } from 'react';
import MainMenu from '../main-menu/components/main-menu';
import InfoName from './util/info-name';
import Description from './util/description';
import TrackButtonContainer from '../containers/untracked-info/track-button';
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
  },
  ticker: {
    paddingBottom: length.mini,
    color: color.blue
  },
  buttonGroup: {
    marginTop: length.medium
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
            <p style={styles.ticker}>{tickerName}</p>
            <h3 style={styles.text}>Description</h3>
            <Description text={description} />
            <div style={styles.buttonGroup}>
              <TrackButtonContainer {...this.props}/>
            </div>
          </div>
        </div>
      </div>
    )
  }
}
