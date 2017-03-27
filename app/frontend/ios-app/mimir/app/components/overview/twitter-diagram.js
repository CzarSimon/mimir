'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { round } from '../../methods/helper-methods';
import { color, length, font } from '../../styles/styles';

export default class TwitterDiagram extends Component {
  constructor(props) {
    super(props);
    this.state = {
      showVolChange: false
    }
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.data.volume !== this.props.data.volume) {
      this.setState({
        showVolChange: true
      });
      setTimeout(() => {
        this.setState({
          showVolChange: false
        });
      }, 500);
    }
  }

  render() {
    const { volume, mean, stdev, minute } = this.props.data;
    const damping = parseFloat(minute) / 60.0;
    const volumeStyle = (this.state.showVolChange) ? styles.volumeChange : styles.text
    return (
      <View style={styles.container}>
        <Text style={volumeStyle}>Tweets in the last hour: {volume}</Text>
        <Text style={styles.text}>Mean: {round(damping * mean, 0)} | Standard Deviation: {round(damping * stdev, 0)}</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    marginVertical: length.small,
    marginHorizontal: length.medium,
    marginBottom: length.small,
    backgroundColor: color.white,
    borderWidth: 1,
    borderColor: color.grey.background,
    padding: length.small
  },
  volumeChange: {
    color: color.red,
    fontFamily: font.type.sans.bold,
    fontSize: font.text
  },
  text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  }
})
