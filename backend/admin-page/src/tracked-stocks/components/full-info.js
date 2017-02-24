import React, { Component } from 'react';
import Description from '../../components/util/description';
import ButtonControlsContainer from '../containers/button-controls';
import { length, color } from '../../styles/styles';

const styles = {
  fullInfo: {
    marginTop: length.small
  },
  siteLink: {
    marginBottom: length.small
  },
  link: {
    color: color.blue
  }
}

export default class FullInfo extends Component {
  render() {
    const { Ticker, Description: desc, Website } = this.props
    return (
      <div style={styles.fullInfo}>
        <p style={styles.siteLink}>
          Website: <a style={styles.link} href={Website}>{Website}</a>
        </p>
        <Description text={desc} />
        <ButtonControlsContainer Ticker={Ticker} />
      </div>
    )
  }
}
