import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, font, length } from '../../../styles/styles';
import { createSubjectString } from '../../../methods/helper-methods';

export default class Info extends Component {
  render() {
    const { compound_score, twitter_references, timestamp } = this.props;
    return (
      <View style={styles.container}>
        <Text style={styles.subject_line}>Subjects: {createSubjectString(compound_score)}</Text>
        <View style={styles.last_row}>
          <Text style={styles.text}>Tweet References: {twitter_references.length}</Text>
          <Text style={styles.text}>{timestamp}</Text>
        </View>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center',
    marginBottom: length.mini
  },
  subject_line: {
    fontSize: font.text,
    fontFamily: font.type.sans.normal,
    color: color.black,
    opacity: 0.9
  },
  last_row: {
    flexDirection: 'row',
    marginTop: length.mini,
    justifyContent: 'space-between'
  },
  text: {
    fontSize: font.text,
    fontFamily: font.type.sans.normal,
    color: color.black,
    opacity: 0.9
  }
});
