import React, { Component } from 'react';
import MenuItem from './main-menu/menu-item';

export default class MainMenu extends Component {
  render() {
    return (
      <div className="main-menu">
        <MenuItem idName={"home"} path={"/"} name={"admin page | "} />
        <MenuItem idName={"tracked-stocks"} path={"/tracked-stocks"} name={"tracked stocks | "} />
        <MenuItem idName={"untracked-tickers"} path={"/untracked-tickers"} name={"untracked tickers | "} />
        <a id="mimir" href="http://mimirapp.co/">mimir</a>
      </div>
    )
  }
}
