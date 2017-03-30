import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import Separator from '../helpers/separator'
import { length, font, color } from '../../styles/styles'
import { card } from '../../styles/common'

export default class Topics extends Component {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.titleGroup}>
          <Text style={styles.title}>Topics</Text>
          <Separator />
        </View>
        <View style={styles.topics}>
          <View style={styles.topicCard}>
            <Text style={styles.topicText}>Comming soon...</Text>
          </View>
        </View>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    marginVertical: length.small
  },
  topics: {
    marginVertical: length.small
  },
  topicText: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    padding: length.small
  },
  topicCard: {
    marginHorizontal: length.medium,
    marginBottom: length.small,
    ...card
  },
  title: {
    fontFamily: font.type.sans.bold,
    fontSize: font.h5,
    marginBottom: length.small
  },
  titleGroup: {
    marginLeft: length.medium
  }
})
