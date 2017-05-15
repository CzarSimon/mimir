import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import { color, font, length } from '../../../styles/styles'
import ArticleButton from './article-button'

export default class ArticleSummary extends Component {
  render() {
    const { summary, url } = this.props
    return (
      <View style={styles.container}>
        <Text style={styles.text}>{summary}</Text>
        <ArticleButton url={url} />
      </View>
    )
  }
}


const styles = StyleSheet.create({
  container: {
    flex: 1
  },
  text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    color: color.black,
    opacity: 0.9,
    paddingBottom: length.small
  }
});
