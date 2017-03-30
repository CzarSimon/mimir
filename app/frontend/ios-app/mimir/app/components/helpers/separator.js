import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native';
import { color, length } from '../../styles/styles';

export default class Separator extends Component {
  render() {
    const styles = createStyles(this.props.customStyles)
    return <View style={styles.separator}/>
  }
}

const createStyles = customStyles => StyleSheet.create({
  separator: {
    alignSelf: 'stretch',
    backgroundColor: color.blue,
    height: 1,
    justifyContent: 'center',
    ...customStyles
  }
})
