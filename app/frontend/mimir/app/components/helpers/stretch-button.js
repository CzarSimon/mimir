import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { color, font, length } from '../../styles/styles';

export default class StretchButton extends Component {
  render() {
    const { btn_color, text } = this.props;
    return (
      <View style={styles.container}>
        <Text style={styles.text}>{text}</Text>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1
  },
  text: {
    fontFamily: font.type.sans.normal,
    fontSize: font.text,
    marginVertical: length.small
  }
});
