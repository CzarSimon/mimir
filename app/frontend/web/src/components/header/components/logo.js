import React, { Component } from 'react';
import { browserHistory } from 'react-router';
import logo from '../../../assets/images/mimir-white.svg';
import { length } from '../../../styles/main';

const style = {
  logo: {
    height: length.icons.medium,
    flex: 1
  }
}

export default class Logo extends Component {
  goToHome = () => {
    browserHistory.push('/');
  }

  render() {
    return (
      <img
        src={logo}
        style={style.logo}
        alt="mimir logo"
        onClick={this.goToHome} />
    )
  }
}
