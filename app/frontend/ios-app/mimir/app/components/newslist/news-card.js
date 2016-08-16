'use strict';

import React, { Component } from 'react';
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native';
import { color, font, length } from '../../styles/styles';
import { create_subject_string } from '../../methods/helper-methods';
const _ = require('lodash');

export default class NewsCard extends Component {
  handle_click = (title) => {
    const { url } = this.props.article_info;
    console.log(`Going to URL: ${url}`);
  }
  render() {
    const { title, compound_score, timestamp, twitter_references } = this.props.article_info;
    const clean_title = _clean_title(title);
    return (
      <TouchableHighlight
        onPress = {() => this.handle_click(clean_title)}
        underlayColor = {color.grey.background}>
        <View style={styles.card}>
          <Text style={styles.title}>{clean_title}</Text>
          <Text style={styles.subject_line}>Subjects: {create_subject_string(compound_score)}</Text>
          <View style={styles.last_row}>
            <Text style={styles.text}>Tweet References: {twitter_references.length}</Text>
            <Text style={styles.text}>{timestamp}</Text>
          </View>
        </View>
      </TouchableHighlight>
    );
  }
}

const _clean_title = (title) => {
  const forbidden = ['TickerLens'];
  const split_title = _.split(_.trim(title), '-');
  const formated_title = (split_title.length < 2) ? split_title : _.join(_.initial(split_title), '-');
  return _.reduce(forbidden, (prev, forb) => _.replace(prev, forb, ""), formated_title);
}

const styles = StyleSheet.create({
  card: {
    flex: 1,
    alignSelf: 'stretch',
    padding: length.small,
    marginHorizontal: length.medium,
    marginBottom: length.small,
    borderColor: color.grey.background,
    borderWidth: 1,
    borderBottomWidth: 2
  },
  title: {
    fontSize: font.h4,
    fontFamily: font.type.sans.normal,
    color: color.blue
  },
  subject_line: {
    fontSize: font.text,
    fontFamily: font.type.sans.normal,
    marginTop: length.mini,
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
})
