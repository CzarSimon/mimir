import React, { Component } from 'react';
import Button from '../../components/util/button';
import { length, color } from '../../styles/styles';
import { portraitMode } from '../../methods/helper-methods';

const screenAdjustment = {
  marginBottom: (!portraitMode()) ? "0" : length.small,
  width: (!portraitMode()) ? "33%" : "100%"
}

const styles = {
  nonSpam: {
    backgroundColor: color.green,
    clear: 'both',
    float: 'left',
    ...screenAdjustment
  },
  spam: {
    paddingLeft: length.large,
    paddingRight: length.large,
    backgroundColor: color.red,
    ...screenAdjustment
  },
  skip: {
    paddingLeft: length.large,
    paddingRight: length.large,
    backgroundColor: color.grey.light,
    color: color.black,
    ...screenAdjustment
  }
}

export default class ButtonGroup extends Component {
  labelSpam = label => {
    const { candidate, labelTweet } = this.props
    labelTweet(candidate, label);
  }

  skip = () => {
    console.log("skipped");
    this.props.skip()
  }

  render() {
    return (
      <div>
        <Button
          text={'NON-SPAM'}
          customStyles={styles.nonSpam}
          handleClick={() => this.labelSpam('NON-SPAM')}
        />
        <Button
          text={'SPAM'}
          customStyles={styles.spam}
          handleClick={() => this.labelSpam('SPAM')}
        />
        <Button
          text={'Skip'}
          customStyles={styles.skip}
          handleClick={this.skip}
        />
      </div>
    )
  }
}
