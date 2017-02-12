import React, { Component } from 'react';
import Description from '../../components/util/description';
import ButtonControlsContainer from '../containers/button-controls';
import { length } from '../../styles/styles';

const styles = {
  fullInfo: {
    marginTop: length.small
  }
}

export default class FullInfo extends Component {
  render() {
    const { Ticker, Description: desc } = this.props
    return (
      <div style={styles.fullInfo}>
        <Description text={desc} />
        <ButtonControlsContainer Ticker={Ticker} />
      </div>
    )
  }
}
