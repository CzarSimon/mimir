import React, { Component } from 'react';
import { color, length, font } from '../../../styles/main';

const style = {
  buttonGroup: {
    flex: 3,
    display: 'flex',
    color: color.blue,
    alignItems: 'center'
  },
  button: {
    flex: 1,
    border: 'solid',
    borderWidth: '1px',
    borderColor: color.blue,
    borderRadius: 0,
    textAlign: 'center'
  },
  buttonLeft: {
    borderRightWidth: '1px',
    borderTopLeftRadius: length.mini,
    borderBottomLeftRadius: length.mini
  },
  buttonRight: {
    borderLeftWidth: '1px',
    borderTopRightRadius: length.mini,
    borderBottomRightRadius: length.mini
  },
  buttonMiddle: {
    borderLeftWidth: 0,
    borderRightWidth: 0
  },
  selectedButton: {
    backgroundColor: color.blue,
    color: color.white
  },
  text: {
    fontSize: font.size.small,
    marginBottom: length.mini,
    marginTop: length.mini
  }
}

export default class PeriodButtons extends Component {
  handleClick = period => {
    this.props.select(period);
  }

  buttonStyle = (period, typeStyle) => {
    const { button, selectedButton } = style;
    if (period === this.props.period) {
      return {
        ...button,
        ...typeStyle,
        ...selectedButton
      }
    } else {
      return {
        ...button,
        ...typeStyle
      }
    }
  }

  render() {
    const { buttonGroup, buttonLeft, buttonRight, text } = style;
    return (
      <div style={buttonGroup}>
        <div
          style={this.buttonStyle('TODAY', buttonLeft)}
          className="period-button"
          onClick={() => this.handleClick('TODAY')}>
          <p style={text}>Today</p>
        </div>
        <div
          className="period-button"
          style={this.buttonStyle('1W', style.buttonMiddle)}
          onClick={() => this.handleClick('1W')}>
          <p style={style.text}>1W</p>
        </div>
        <div
          className="period-button"
          style={this.buttonStyle('1M', buttonRight)}
          onClick={() => this.handleClick('1M')}>
          <p style={style.text}>1M</p>
        </div>
      </div>
    );
  }
}
