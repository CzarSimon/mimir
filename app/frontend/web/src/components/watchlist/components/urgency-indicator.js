import React, { Component } from 'react';
import { length, color, font } from '../../../styles/main';
import { classifyUrgency, LOW_LABEL, HIGH_LABEL, URGENT_LABEL } from '../../../methods/server/twitter-miner';

const style = {
  urgent: {
    display: 'table-cell',
    width: length.button,
    marginRight: '-10px',
    fontSize: font.size.large,
    fontWeight: 'bold',
    textAlign: 'center',
    float: 'left',
    height: '100%'
  },
  nonUrgent: {
    height: '0',
    width: '0'
  }
}

export default class UrgencyInicator extends Component {
  shouldComponentUpdate(nextProps) {
    return nextProps.volume !== this.props.volume;
  }

  getStyle = urgencyLevel => {
    switch (urgencyLevel) {
      case HIGH_LABEL:
        return {
          ...style.urgent,
          color: color.yellow
        };
      case URGENT_LABEL:
        return {
          ...style.urgent,
          color: color.red
        };
      default:
        return style.nonUrgent
    }
  }

  render() {
    const urgencyLevel = classifyUrgency(this.props);
    if (urgencyLevel === LOW_LABEL) {
      return <div />
    } else {
      return (
        <div style={this.getStyle(urgencyLevel)}>
          <p>!</p>
        </div>
      )
    }
  }
}
