import React, { Component } from 'react';
import IoIosSearch from 'react-icons/lib/io/ios-search';
import { color, length, font } from '../../../styles/main';

const style = {
  icon: {
    flex: 1,
    marginTop: '0',
    color: color.white,
    marginRight: length.medium
  },
  text: {
    flex: 1,
    marginTop: '0',
    paddingTop: length.mini,
    fontSize: font.size.small,
    marginRight: length.large
  }
}

export default class SearchButton extends Component {
  render() {
    const { active, goToSearch, keyboardDown, cancelSearch } = this.props;
    if (!active) {
      return <IoIosSearch onClick={goToSearch} size={35} style={style.icon} />
    } else if (!keyboardDown) {
      return <p style={style.text} onClick={cancelSearch}>Cancel</p>
    } else {
      return <div />
    }
  }
}
