'use strict';
import React, { Component } from 'react';
import { View, Text, StyleSheet, TabBarIOS } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { select_tab } from '../actions/navigation.actions';

class CompanyPage extends Component {
  render() {
    console.log(this.props);
    return (
      <View style={styles.container}>
        <Text>Company Page</Text>
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
)(CompanyPage);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center'
  }
})
