import React, { Component } from 'react';
import { browserHistory } from 'react-router';
import { length, color, font } from '../../../styles/main';
import IoIosPlusOutline from 'react-icons/lib/io/ios-plus-outline';

const style = {
  container: {
    padding: length.small,
    paddingLeft: length.small,
    marginBottom: length.small,
    display: 'flex'
  },
  name: {
    fontSize: font.size.medium,
    margin: '0',
    marginBottom: length.small,
    color: color.black
  },
  ticker: {
    fontSize: font.size.small,
    margin: '0',
    color: color.grey.dark
  },
  info: {
    flex: 4,
    padding: '0'
  },
  addButton: {
    flex: 1,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center'
  },
  icon: {
    color: color.green
  }
}

export default class Search extends Component {
  addTicker = () => {
    const { ticker, userId, addNewTicker } = this.props;
    console.log('ticker:', ticker);
    console.log('userId:', userId);
    addNewTicker(userId, ticker);
    browserHistory.push("/");
  }

  render() {
    const { name, ticker } = this.props;
    return (
      <div className="card" style={style.container}>
        <div style={style.info}>
          <p style={style.name}>{name}</p>
          <p style={style.ticker}>{ticker}</p>
        </div>
        <div style={style.addButton} onClick={this.addTicker}>
          <IoIosPlusOutline size={35} style={style.icon}/>
        </div>
      </div>
    )
  }
}
