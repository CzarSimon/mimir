import React, { Component } from 'react';
import { font, color, length } from '../../../styles/main';
import NewsCard from './news-card';
import BasicInfoContainer from '../../basic-info/containers/main';

const style = {
  container: {
    display: 'flex',
  },
  header: {
    flex: 5,
    color: color.blue,
    marginTop: length.small,
    fontSize: font.size.large,
    marginBottom: length.small,
    paddingLeft: length.small
  }
}

export default class Newslist extends Component {
  renderItem = (newsInfo, key) => {
    return <NewsCard key={key} {...newsInfo} />
  }

  render() {
    const { news } = this.props;
    return (
      <div>
        <BasicInfoContainer />
        <div style={style.container}>
          <p style={style.header}>Top News</p>
        </div>
        {news.map(this.renderItem)}
      </div>
    )
  }
}
