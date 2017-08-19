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
  handleTabClick = (clickedTab) => {
    const { selectedTab, handleClick } = this.props;
    if (clickedTab !== selectedTab) {
      handleClick(clickedTab);
    }
  }
  render() {
    const { company, twitterData, selectedTab, handleClick } = this.props;
    const tabs = {
      news: <NewsContainer company={company} />,
      overview: <OverviewContainer company={company} twitterData={twitterData}/>
    };
    const iconNames = {
      news: 'ios-paper-outline',
      overview: 'ios-information-circle-outline'
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
              iconName={iconNames[name]}
              iconSize={length.icons.medium}
              selected={selectedTab === name}
              onPress={() => this.handleTabClick(name)}>
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
