import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Separator from '../helpers/separator';
import { length } from '../../styles/styles';

export default class Topics extends Component {
  render() {
    return (
      <View style={styles.container}>
        <Text>Topics</Text>
        <Separator />
        <View style={styles.topics}>
          <Text>Comming soon...</Text>
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
  }
})
