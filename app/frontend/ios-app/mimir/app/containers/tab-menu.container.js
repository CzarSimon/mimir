'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import TabMenu from '../components/tab-menu';
import { select_tab } from '../actions/navigation.actions';

class TabMenuContainer extends Component {
  handle_tab_click(clicked_tab) {
    this.props.actions.select_tab(clicked_tab);
  }
  render() {
    const { selected_tab } = this.props.state.navigation;
    return (
      <View style={styles.container}>
        <TabMenu
          nav={this.props.navigator}
          selected_tab={selected_tab}
          handle_click={this.handle_tab_click.bind(this)}
        />
      </View>
    );
  }
}

export default connect(
  (state) => ({ state: { navigation: state.navigation } }),
  (dispatch) => ({ actions: bindActionCreators({ select_tab }, dispatch) })
)(TabMenuContainer);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center'
  }
})
