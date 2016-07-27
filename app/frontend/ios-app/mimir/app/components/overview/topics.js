import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Separator from '../helpers/separator';
import { length, font } from '../../styles/styles';

export default class Topics extends Component {
  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.title}>Topics</Text>
        <Separator />
        <View style={styles.topics}>
          <Text style={styles.topic_text}>Comming soon...</Text>
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
    fontSize: font.text
  },
  title: {
    fontFamily: font.type.sans.bold,
    fontSize: font.h5,
    marginBottom: length.mini
  }
})
