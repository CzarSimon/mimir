import React, { Component } from 'react';
import { View, Text, TouchableHighlight, StyleSheet } from 'react-native';
import { length, color, font } from '../../../styles/styles';

export default class LogoutButton extends Component {
  render() {
    return (
      <View style={style.container}>
        <TouchableHighlight onPress={this.props.logout}>
          <View>
            <Text>Logout</Text>
          </View>
        </TouchableHighlight>
      </View>
    );
  }
}

const style = StyleSheet.create({
  container: {
    flex: 1,
    borderWidth: 1
  }
})
