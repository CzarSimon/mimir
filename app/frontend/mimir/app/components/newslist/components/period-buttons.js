import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import _ from 'lodash';
import { ButtonGroup } from 'react-native-elements';
import { color, length, font } from '../../../styles/styles';

export default class PeriodButtons extends Component {
  constructor(props) {
    super(props);
    this.periods = ['Today', '1W', '1M']
  }

  handleClick = periodIndex => {
    const newPeriod = _.toUpper(this.periods[periodIndex]);
    this.props.select(newPeriod);
  }

  selectedIndex = () => {
    return _.findIndex(this.periods, p => _.toUpper(p) === this.props.period)
  };

  render() {
    return (
        <ButtonGroup
          onPress={this.handleClick}
          selectedIndex={this.selectedIndex()}
          buttons={this.periods}
          textStyle={style.text}
          selectedTextStyle={style.selectedText}
          underlayColor={color.blue}
          buttonStyle={style.button}
          innerBorderStyle={{color: color.blue}}
          selectedBackgroundColor={color.blue}
          containerStyle={style.buttonGroup} />
    );
  }
}

const style = StyleSheet.create({
  buttonGroup: {
    flex: 1,
    backgroundColor: color.grey.appBackground,
    borderWidth: 1,
    borderColor: color.blue,
    borderRadius: length.mini
  },
  button: {
    margin: 0
  },
  text: {
    fontSize: font.h5,
    color: color.blue
  },
  selectedText: {
    ...this.text,
    color: color.white
  }
});
