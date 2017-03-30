'use strict'
import React, { Component } from 'react'
import { View, Text, StyleSheet, TouchableHighlight } from 'react-native'
import { length, color } from '../../styles/styles'
import Icon from 'react-native-vector-icons/Ionicons'
import { last } from 'lodash'
import { SEARCH_PAGE } from '../../routing/main'

export default class BackButton extends Component {
  handleClick = navigator => {
    const lastRoute = last(navigator.getCurrentRoutes())
    console.log(lastRoute);
    if (lastRoute.name === SEARCH_PAGE) {
      console.log(lastRoute);
    }
    navigator.pop()
  }

  render() {
    const { index, handleClick } = this.props
    if (index === 0) {
      return (<View style={styles.container}/>)
    } else {
      return (
        <TouchableHighlight
          onPress={() => handleClick()}
          underlayColor={color.grey.background}>
          <View style={styles.container}>
              <View style={styles.button}>
                <Icon name='ios-arrow-back-outline' size={length.icons.medium} color={color.blue} />
              </View>
          </View>
      </TouchableHighlight>
      )
    }
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    alignSelf: 'stretch',
    justifyContent: 'center',
    width: length.button
  },
  button: {
    flex: 1,
    paddingVertical: length.mini + 2
  }
})
