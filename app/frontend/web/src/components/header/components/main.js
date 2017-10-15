import React, { Component } from 'react';
import { color, length } from '../../../styles/main';
import Logo from './logo';
import Title from './title';
import SearchButtonContainer from '../containers/search-button';
import SearchBarContainer from '../containers/search-bar';

const style = {
  header: {
    height: length.icons.medium,
    top: '0',
    width: '100%',
    backgroundColor: color.blue,
    position: 'fixed',
    padding: length.small,
    color: color.white,
    display: 'flex'
  }
}

const MAIN_ROUTE = '/';
const LOGIN_ROUTE = '/login';
const SEARCH_ROUTE = '/search';

export default class Header extends Component {
  getMiddleComponent = () => {
    const { path, ticker } = this.props;
    switch(path) {
      case MAIN_ROUTE:
        return <Title text='mimir' />
      case SEARCH_ROUTE:
        return <SearchBarContainer />
      case LOGIN_ROUTE:
        return <Title text={""} />
      default:
        return <Title text={ticker} />
    }
  }

  render() {
    const { path } = this.props;
    return (
      <div style={style.header} className="header">
        <Logo />
        {this.getMiddleComponent()}
        {(path !== LOGIN_ROUTE) ? <SearchButtonContainer /> : <div />}
      </div>
    )
  }
}
