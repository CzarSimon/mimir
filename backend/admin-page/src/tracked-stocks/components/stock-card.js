import React, { Component } from 'react';
import FullInfo from './full-info';
import { length, font, color } from '../../styles/styles';

const styles = {
  card: {
    marginBottom: length.medium,
    padding: length.medium,
    backgroundColor: color.white,
    marginRight: '4vw'
  },
  name: {
    fontSize: font.size.large,
    marginBottom: length.mini
  },
  ticker: {
    fontSize: font.size.small,
    color: color.blue
  }
}

export default class StockCard extends Component {
  constructor(props) {
    super(props)
    this.state = {
      clicked: false
    }
  }

  handleClick = () => {
    this.setState({
      clicked: !this.state.clicked
    })
  }

  render() {
    const { Name, Ticker } = this.props;
    const fullInfo = (this.state.clicked) ? <FullInfo {...this.props} /> : <div />
    return (
      <div className='card' style={styles.card}>
        <div onClick={this.handleClick}>
          <p style={styles.name}>{Name}</p>
          <p style={styles.ticker}>{Ticker}</p>
        </div>
        {fullInfo}
      </div>
    )
  }
}
