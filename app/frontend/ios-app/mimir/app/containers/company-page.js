'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TabBarIOS } from 'react-native';

export default class CompanyPage extends Component {
  render() {
    console.log(this.props);
    return (
      <View style={styles.container}>
        <Text>Company Page</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center'
  }
})
