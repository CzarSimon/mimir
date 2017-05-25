import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { head } from 'lodash';
import { labelTweet, fetchSpamCandidates, skipCandidate } from '../../reducers/spam';
import Spam from '../components/main';


class SpamContainer extends Component {
  componentDidMount() {
    const { state, actions } = this.props;
    actions.fetchSpamCandidates(state.token);
  }

  labelCandidate = (tweet, label) => {
    const { state, actions } = this.props;
    const labeledTweet = {
      Tweet: tweet,
      Label: label
    }
    actions.labelTweet(labeledTweet, state.token);
  }

  skip = () => {
    this.props.actions.skipCandidate();
  }

  render() {
    const { labeledCount, candidates } = this.props.state;
    const loaded = (candidates.length > 0);
    return (
      <Spam
        labelTweet={this.labelCandidate}
        skip={this.skip}
        count={labeledCount}
        loaded={loaded}
        candidate={(loaded) ? head(candidates) : undefined}
      />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    ...state.spam,
    token: state.user.token
  }
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators({
    labelTweet,
    fetchSpamCandidates,
    skipCandidate
  }, dispatch)
})

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToProps(dispatch)
)(SpamContainer);
