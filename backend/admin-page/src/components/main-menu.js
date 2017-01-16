import React, { Component } from 'react';
import MenuItem from './main-menu/menu-item';
import { color, length } from '../styles/styles';

const styles = {
  menu: {
    height: '100vh',
    width: '20vw',
    float: 'left',
    backgroundColor: color.blue,
    color: color.white,
    marginRight: length.large
  },
  list: {
    marginTop: length.large,
    padding: '0',
    listStyleType: 'none',
    textAlign: 'center'
  }
}

export default class MainMenu extends Component {
  render() {
    return (
      <div className="main-menu" style={styles.menu}>
        <ul style={styles.list}>
          <MenuItem idName={"home"} path={"/"} name={"admin page"} />
          <MenuItem idName={"tracked-stocks"} path={"/tracked-stocks"} name={"tracked stocks"} />
          <MenuItem idName={"untracked-tickers"} path={"/untracked-tickers"} name={"untracked tickers"} />
          <li><a id="mimir" href="http://mimirapp.co/">mimir</a></li>
        </ul>
      </div>
    )
  }
}
