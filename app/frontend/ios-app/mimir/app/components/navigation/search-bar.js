'use strict';

import React, { Component } from 'react';
import { View, Text, TextInput, StyleSheet } from 'react-native';
import { trim } from 'lodash';
import { length, color, font } from '../../styles/styles';

export default class SearchBar extends Component {
  constructor(props) {
    super(props);
    this.state = {
      text: "Search ticker..."
    };
  }

  handle_submit = () => {
    this.props.run_query(this.state.text);
  }

  handle_new_text = (new_text) => {
    this.setState({
      text: trim(new_text)
    })
  }

  render() {
    return (
      <View style={styles.container}>
        <TextInput
          style={styles.search_box}
          onChangeText={(text) => this.handle_new_text(text)}
          value={"   " + this.state.text}
          selectionColor={color.green}
          clearButtonMode='always'
          returnKeyType='search'
          autoCorrect={false}
          onSubmitEditing={() => this.handle_submit()}
          onFocus={() => this.setState({text: ""})}
        />
      </View>
      );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: 'stretch',
    marginHorizontal: length.button
  },
  search_box: {
    flex: 1,
    margin: length.mini + 3,
    borderRadius: 3,
    backgroundColor: color.grey.background,
    color: color.green,
    fontFamily: font.type.sans.normal,
    fontSize: font.text
  }
})
