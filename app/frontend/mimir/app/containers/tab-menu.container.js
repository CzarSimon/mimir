'use strict'
import React, { Component } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import TabMenu from '../components/tab-menu';
import BasicInfo from '../components/tab-menu/basic-info';
import { selectTab } from '../ducks/navigation';
import { arr_get_value_by_key } from '../methods/helper-methods';

class TabMenuContainer extends Component {
  handleTabClick = clickedTab => {
    this.props.actions.selectTab(clickedTab);
  }

  render() {
    const { stocks, navigation } = this.props.state;
    const { selectedTab, activeTicker } = navigation;
    const company = stocks.data[activeTicker];
    const twitterData = this.props.state.twitterData.data[activeTicker];

    return (
      <View style={styles.container}>
        <BasicInfo company={company} twitterData={twitterData} />
        <TabMenu
          company = {company}
          twitterData = {twitterData}
          selectedTab = {selectedTab}
          handleClick = {this.handleTabClick}
        />
      </View>
    )
  }
}

const mapStateToProps = state => ({
  state: {
    twitterData: state.twitterData,
    stocks: state.stocks,
    navigation: state.navigation
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    selectTab
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(TabMenuContainer)

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'stretch',
    justifyContent: 'center'
  }
})
