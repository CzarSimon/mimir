'use strict'

import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, font, length } from '../../styles/styles'
import { card } from '../../styles/common'
import { createCleanTitle } from '../../methods/helper-methods'
import Summary from './news-card/summary'
import Info from './news-card/info'
import ArticleButton from './news-card/article-button'
import SafariView from 'react-native-safari-view'

export default class NewsCard extends Component {
  constructor(props) {
    super(props)
    this.state = {
      clicked: false
    }
  }

  handleClick = () => {
    this.setState({
      clicked: !this.state.clicked
    })
  }

  summaryComponent = () => {
    const { summary, twitter_references, timestamp, compound_score, url } = this.props.article_info
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

  render() {
    const { title } = this.props.article_info
    const cleanTitle = createCleanTitle(title)
    return (
      <TouchableHighlight
        onPress = {() => this.handleClick()}
        underlayColor = {color.grey.background}>
        <View style={styles.card}>
          <Text style={styles.title}>{cleanTitle}</Text>
          {this.summaryComponent()}
        </View>
      </TouchableHighlight>
    )
  }
}

const styles = StyleSheet.create({
  card: {
    flex: 1,
    alignItems: 'stretch',
    padding: length.small,
    marginHorizontal: length.medium,
    marginBottom: length.small,
    ...card
  },
  title: {
    fontSize: font.h4,
    fontFamily: font.type.sans.bold,
    color: color.black,
    marginBottom: length.mini
  }
})
