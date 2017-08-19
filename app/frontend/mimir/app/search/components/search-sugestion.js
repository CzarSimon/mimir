'use strict'

import React, { Component } from 'react';
import { length, color, font } from '../../styles/styles';
import { View, Text, StyleSheet, TouchableWithoutFeedback } from 'react-native';

export default class SearchSugestion extends Component {
  handleClick = () => {
    const { sugestion, updateAndRunQuery } = this.props;
    updateAndRunQuery(sugestion.name);
  }

  render() {
    const { name } = this.props.sugestion;
    return (
      <View style={styles.sugestion}>
        <TouchableWithoutFeedback onPress={this.handleClick}>
          <View>
            <Text style={styles.sugestionText}>
              {name}
            </Text>
          </View>
        </TouchableWithoutFeedback>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  sugestion: {
    alignItems: 'center'
  },
  sugestionText: {
    color: color.grey.dark,
    fontFamily: font.type.sans.normal,
    fontSize: font.h5,
    paddingVertical: length.small
  }
});
