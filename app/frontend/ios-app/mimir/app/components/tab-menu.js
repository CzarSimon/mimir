'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TabBarIOS } from 'react-native';
import { map, capitalize } from 'lodash';
import { color } from '../styles/styles';

export default class TabMenu extends Component {
  handle_tab_click = (clicked_tab) => {
    const { selected_tab, handle_click } = this.props;
    if (clicked_tab !== selected_tab) {
      handle_click(clicked_tab);
    }
  }
  render() {
    const tabs = {
      overview: <View style={styles.container}><Text>Overview</Text></View>,
      news: <View style={styles.container}><Text>News</Text></View>,
      tweets: <View style={styles.container}><Text>Tweets</Text></View>,
      statistics: <View style={styles.container}><Text>Statistics</Text></View>
    }
    const { selected_tab, handle_click } = this.props;
    return (
      <TabBarIOS
        style={styles.tab_bar}
        tintColor={color.green}>
        {
          map(tabs, (component, name) => (
            <TabBarIOS.Item
              key={name}
              title={capitalize(name)}
              selected={selected_tab === name}
              onPress={() => this.handle_tab_click(name)}>
              {component}
            </TabBarIOS.Item>
          ))
        }
      </TabBarIOS>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center'
  },
  tab_bar: {
    borderTopWidth: 4,
    borderTopColor: color.red,
  }
})
