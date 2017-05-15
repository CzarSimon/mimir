'use strict'

import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { color, font, length } from '../../styles/styles'
import { card } from '../../styles/common'
import { createCleanTitle } from '../../methods/helper-methods'
import ArticleSummary from './news-card/summary'
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
    const { Summary, Twitter_References, Timestamp, Compound_Score, URL } = this.props.articleInfo
    if (this.state.clicked) {
      return <ArticleSummary url={URL} summary={Summary} />
    } else {
      return (
        <Info
          twitter_references={Twitter_References}
          compound_score={Compound_Score}
          timestamp={Timestamp}
        />
      )
    }
  }

  render() {
    const { Title } = this.props.articleInfo
    const cleanTitle = createCleanTitle(Title)
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
