'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { color, font, length } from '../../styles/styles';
import { create_subject_string, create_clean_title } from '../../methods/helper-methods';
import Summary from './news-card/summary';
import Info from './news-card/info';
import ArticleButton from './news-card/article-button';
import SafariView from 'react-native-safari-view';

export default class NewsCard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      clicked: false
    }
  }

  handle_click = () => {
    this.setState({
      clicked: !this.state.clicked
    })
  }

  summary_component = () => {
    const { summary, twitter_references, timestamp, compound_score, url } = this.props.article_info;
    if (this.state.clicked) {
      return <Summary url={url} summary={summary} />
    } else {
      return (
        <Info
          twitter_references={twitter_references}
          compound_score={compound_score}
          timestamp={timestamp}
        />
      )
    }
  }

  go_to_article = () => {
    console.log("Clicked go to article button")
  }

  render() {
    const { title, compound_score, timestamp, twitter_references, summary, url } = this.props.article_info;
    const clean_title = create_clean_title(title);
    return (
      <TouchableHighlight
        onPress = {() => this.handle_click()}
        underlayColor = {color.grey.background}>
        <View style={styles.card}>
          <Text style={styles.title}>{clean_title}</Text>
          {this.summary_component()}
        </View>
      </TouchableHighlight>
    );
  }
}

const styles = StyleSheet.create({
  card: {
    flex: 1,
    alignItems: 'stretch',
    padding: length.small,
    marginHorizontal: length.medium,
    marginBottom: length.small,
    borderColor: color.grey.background,
    borderWidth: 1,
    borderBottomWidth: 2,
    backgroundColor: color.white
  },
  title: {
    fontSize: font.h4,
    fontFamily: font.type.sans.bold,
    color: color.black,
    marginBottom: length.mini
  }
})
