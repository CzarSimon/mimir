'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import TabMenu from '../components/tab-menu';
import BasicInfo from '../components/tab-menu/basic-info';
import { select_tab } from '../actions/navigation.actions';
import { article_route } from '../routing/routes';
import { arr_get_value_by_key } from '../methods/helper-methods';

class TabMenuContainer extends Component {
  handle_tab_click(clicked_tab) {
    this.props.actions.select_tab(clicked_tab);
  }

  render() {
    const { user, stocks, navigation } = this.props.state;
    const { selected_tab, active_ticker } = navigation;
    const company = stocks.data[active_ticker]
    const twitter_data = user.twitter_data.data[active_ticker];

    return (
      <View style={styles.container}>
        <BasicInfo
          company = {company}
          twitter_data = {twitter_data}
        />
        <TabMenu
          company = {company}
          twitter_data = {twitter_data}
          selected_tab = {selected_tab}
          handle_click = {this.handle_tab_click.bind(this)}
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
      select_tab
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
