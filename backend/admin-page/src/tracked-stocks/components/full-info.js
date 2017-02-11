import React, { Component } from 'react';
import Description from '../../components/util/description';
import ButtonControls from './button-controls';
import { length } from '../../styles/styles';

const styles = {
  fullInfo: {
    marginTop: length.small
  }
}

export default class FullInfo extends Component {
  render() {
    const { Description: desc } = this.props
    return (
      <div style={styles.fullInfo}>
        <Description text={desc}/>
        <ButtonControls />
      </div>
    )
  }
}
