import React, { Component } from 'react';
import { color, length } from '../../../styles/main';
import Logo from './logo';
import Title from './title';
import SearchButtonContainer from '../containers/search-button';
import SearchBarContainer from '../containers/search-bar';

const style = {
  header: {
    height: length.icons.medium,
    backgroundColor: color.blue,
    padding: length.small,
    color: color.white,
    marginBottom: length.small,
    display: 'flex'
  }
}

const MAIN_ROUTE = '/';
const SEARCH_ROUTE = '/search';

export default class Header extends Component {
  getMiddleComponent = () => {
    const { path, ticker } = this.props;
    switch(path) {
      case MAIN_ROUTE:
        return <Title text='mimir' />
      case SEARCH_ROUTE:
        return <SearchBarContainer />
      default:
        return <Title text={ticker} />
    }

  }
  render() {
    return (
      <div style={style.header}>
        <Logo />
        {this.getMiddleComponent()}
        <SearchButtonContainer />
      </div>
    )
  }
}
