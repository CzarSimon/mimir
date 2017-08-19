'use strict'

import React, { Component } from 'react';
import { length, font, color } from '../../styles/styles';
import {
  View,
  Text,
  TouchableWithoutFeedback,
  StyleSheet
} from 'react-native';

export default class SearchCancel extends Component {
  render() {
    return (
      <View style={styles.container}>
        <TouchableWithoutFeedback onPress={() => this.props.cancelSearch()}>
          <View style={styles.button}>
            <Text style={styles.buttonText}>Cancel</Text>
          </View>
        </TouchableWithoutFeedback>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    width: length.navebar
  },
  button: {
    paddingRight: length.small
  },
  buttonText: {
    fontFamily: font.type.sans.normal,
    color: color.grey.dark,
    fontSize: font.h5
  }
})
