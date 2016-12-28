import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Separator from '../helpers/separator';
import { length, font, color } from '../../styles/styles';

export default class Topics extends Component {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.title_group}>
          <Text style={styles.title}>Topics</Text>
          <Separator />
        </View>
        <View style={styles.topics}>
          <View style={styles.topic_card}>
            <Text style={styles.topic_text}>Comming soon...</Text>
          </View>
        </View>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    marginVertical: length.small
  },
  topics: {
    marginVertical: length.small
  },
  topic_text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    padding: length.small
  },
  topic_card: {
    marginHorizontal: length.medium,
    marginBottom: length.small,
    backgroundColor: color.white,
    borderWidth: 1,
    borderColor: color.grey.background
  },
  title: {
    fontFamily: font.type.sans.bold,
    fontSize: font.h5,
    marginBottom: length.small
  },
  title_group: {
    marginLeft: length.medium
  }
})
