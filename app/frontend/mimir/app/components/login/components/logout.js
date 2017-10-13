import React, { Component } from 'react';
import { View, Text, TouchableHighlight, StyleSheet } from 'react-native';
import { length, color, font } from '../../../styles/styles';

export default class LogoutButton extends Component {
  render() {
    return (
      <View style={style.container}>
        <TouchableHighlight underlayColor={color.grey.dark} onPress={this.props.logout}>
          <View style={style.button}>
            <Text style={style.text}>Logout</Text>
          </View>
        </TouchableHighlight>
      </View>
    );
  }
}

const style = StyleSheet.create({
  container: {
    flex: 1,
    marginHorizontal: length.button,
    marginVertical: length.button,
    alignItems: 'stretch'
  },
  button: {
    borderWidth: 1,
    alignItems: 'center',
    borderRadius: length.mini,
    borderColor: color.grey.dark,
    padding: length.large
  },
  text: {
    fontFamily: font.type.sans.bold,
    fontSize: font.h3,
    color: color.grey.dark
  }
})
