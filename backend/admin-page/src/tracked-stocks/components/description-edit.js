import React, { Component } from 'react'
import { length, font } from '../../styles/styles'
import Button from '../../components/util/button'

const styles = {
  container: {
    width: '90%',
    marginBottom: length.small
  },
  input: {
    width: '100%',
    padding: length.small,
    marginBottom: length.small,
    fontSize: font.size.small,
    borderStyle: 'none',
  }
}


export default class DescriptionEdit extends Component {
  constructor(props) {
    super(props)
    this.state = {
      description: props.currentDescription
    }
  }

  handleChange = event => {
    this.setState({
      description: event.target.value
    })
  }

  handleSubmit = event => {
    event.preventDefault()
    const { updateInfo, toggleEdit } = this.props
    const { description } = this.state
    updateInfo(description)
    toggleEdit()
  }

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit} style={styles.container}>
          <textarea
            type='textarea'
            rows='4'
            value={this.state.description}
            onChange={this.handleChange}
            style={styles.input}
            className='card'
            autoFocus
          />
          <Button text='Save' />
        </form>
      </div>
    )
  }
}
