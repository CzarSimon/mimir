'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import TabMenu from '../components/tab-menu';
import BasicInfo from '../components/tab-menu/basic-info';
import { selectTab } from '../ducks/navigation';
import { article_route } from '../routing/routes';
import { arr_get_value_by_key } from '../methods/helper-methods';

class TabMenuContainer extends Component {
  handleTabClick = clickedTab => {
    this.props.actions.selectTab(clickedTab);
  }

  render() {
    const { user, stocks, navigation } = this.props.state;
    const { selectedTab, activeTicker } = navigation;
    const company = stocks.data[activeTicker]
    const twitterData = user.twitterData.data[activeTicker];

    return (
      <View style={styles.container}>
        <BasicInfo
          company = {company}
          twitterData = {twitterData}
        />
        <TabMenu
          company = {company}
          twitterData = {twitterData}
          selectedTab = {selectedTab}
          handleClick = {this.handleTabClick}
        />
      </View>
    );
  }
}

export default connect(
  (state) => ({
    state: {
      user: state.user,
      stocks: state.stocks,
      navigation: state.navigation
    }
  }),
  (dispatch) => ({
    actions: bindActionCreators({
      selectTab
    }, dispatch)
  })
)(TabMenuContainer);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center'
  }
})
