'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TabBarIOS } from 'react-native';
import { map, capitalize } from 'lodash';
import { color, length } from '../styles/styles';
import OverviewContainer from '../containers/overview.container';
import NewsContainer from '../containers/news.container';
import StatisticsContainer from '../containers/statistics.container';
import Icon from 'react-native-vector-icons/Ionicons';

export default class TabMenu extends Component {
  handle_tab_click = (clicked_tab) => {
    const { selected_tab, handle_click } = this.props;
    if (clicked_tab !== selected_tab) {
      handle_click(clicked_tab);
    }
  }
  render() {
    const { company, twitter_data, selected_tab, handle_click, select_article } = this.props;
    const tabs = {
      overview: <OverviewContainer company={company} twitter_data={twitter_data}/>,
      news: <NewsContainer company={company} />,
      statistics: <StatisticsContainer company={company}/>
  };
    const icon_names = {
      overview: 'ios-information-circle-outline',
      news: 'ios-paper-outline',
      statistics: 'ios-stats-outline'
    };
    return (
      <TabBarIOS
        style={styles.tab_bar}
        tintColor={color.blue}>
        {
          map(tabs, (component, name) => (
            <Icon.TabBarItemIOS
              key={name}
              title={capitalize(name)}
              iconName={icon_names[name]}
              iconSize={length.icons.medium}
              selected={selected_tab === name}
              onPress={() => this.handle_tab_click(name)}>
              {component}
            </Icon.TabBarItemIOS>
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
