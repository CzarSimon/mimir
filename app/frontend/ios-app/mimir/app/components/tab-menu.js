'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TabBarIOS } from 'react-native';
import { map, capitalize } from 'lodash';
import { color } from '../styles/styles';
import OverviewContainer from '../containers/overview.container';
import StatisticsContainer from '../containers/statistics.container';

export default class TabMenu extends Component {
  handle_tab_click = (clicked_tab) => {
    const { selected_tab, handle_click } = this.props;
    if (clicked_tab !== selected_tab) {
      handle_click(clicked_tab);
    }
  }
  render() {
    const { company, twitter_data, selected_tab, handle_click } = this.props;
    const tabs = {
      overview: <OverviewContainer company={company} twitter_data={twitter_data}/>,
      news: <View style={styles.container}><Text>News component not started yet...</Text></View>,
      tweets: <View style={styles.container}><Text>Tweet component not started yet...</Text></View>,
      statistics: <StatisticsContainer company={company}/>
    }
    return (
      <TabBarIOS
        style={styles.tab_bar}
        tintColor={color.blue}>
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
