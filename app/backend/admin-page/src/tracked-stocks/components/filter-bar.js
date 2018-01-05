import React, { Component } from 'react';
import { length, font } from '../../styles/styles';

const styles = {
  input: {
    width: '94.7%',
    height: '7vh',
    marginBottom: length.small,
    fontSize: font.size.medium,
    borderStyle: 'none',
    textIndent: length.medium,
    outline: 'none'
  }
}

export default class FilterBar extends Component {
  constructor(props) {
    super(props)
    this.state = {
      filterTerm: ''
    }
  }

  handleInput = event => {
    const { handleFilterInput } = this.props
    const { value: filterInput } = event.target
    this.setState({filterTerm: filterInput})
    handleFilterInput(filterInput)
  }

  render() {
    const { filterTerm } = this.state
    return (
      <input
        type='text'
        style={styles.input}
        value={filterTerm}
        onChange={this.handleInput}
        placeholder='Filter stocks by name'
        className='card'
        autoFocus
      />
    )
  }
}
