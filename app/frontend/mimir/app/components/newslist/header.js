import React, { Component } from 'react'
import { View, Text, StyleSheet } from 'react-native'
import PeriodButtonsContainer from './containers/period-buttons';
import { color, font, length } from '../../styles/styles'

export default class Header extends Component {
  render() {
    return (
      <View style={style.container}>
        <Text style={style.headerText}>Top News</Text>
        <View style={style.buttonGroup}>
          <PeriodButtonsContainer />
        </View>
      </View>
    )
  }
}

const style = StyleSheet.create({
  container: {
    alignItems: 'center',
    flexDirection: 'row',
    marginHorizontal: length.medium,
    marginVertical: length.mini,
    height: length.button
  },
  headerText: {
    flex: 2,
    fontSize: font.h3,
    fontFamily: font.type.sans.normal,
    paddingLeft: length.small,
    color: color.blue
  },
  buttonGroup: {
    flex: 5,
    alignSelf: 'flex-end'
  }
})
