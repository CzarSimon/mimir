import React, { Component } from 'react';
import { createCleanTitle } from '../../../methods/helper-methods';
import { length, font, color } from '../../../styles/main';

import Info from './info';

const style = {
  card: {
    width: '100%',
    clear: 'both',
    overflow: 'auto',
    marginBottom: length.small
  },
  title: {
    fontSize: font.size.medium,
    fontWeight: 'bold',
    color: color.black,
    paddingLeft: length.small,
    paddingRight: length.small,
    marginBottom: '0'
  }
}

export default class NewsCard extends Component {
  handleClick = () => {
    const { url } = this.props;
    window.open(url);
  }

  render() {
    const { title, twitterReferences, timestamp } = this.props;
    const cleanTitle = createCleanTitle(title);
    return (
      <div className="card" style={style.card} onClick={this.handleClick}>
        <p style={style.title}>{cleanTitle}</p>
        <Info
          timestamp={timestamp}
          twitterReferences={twitterReferences} />
      </div>
    )
  }
}
