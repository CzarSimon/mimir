import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, font, length } from '../../../styles/styles';

export default class Summary extends Component {
  summary_component = (text) => {
    return (
      <View style={styles.container}>
        <Text style={styles.text}>{text}</Text>
      </View>
    )
  }

  render() {
    const { summary, clicked } = this.props;
    return (clicked) ? this.summary_component(summary) : <View />
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
    opacity: 0.9
  }
});
