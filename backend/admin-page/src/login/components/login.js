import React, { Component } from 'react';
import { length, color } from '../../styles/styles';

const formComponent = {
  width: '100%',
  height: '6vh',
  marginBottom: length.small,
  fontSize: '1.1em',
  borderStyle: 'none',
}

const styles = {
  form: {
    width: '25vw',
    margin: '25vh auto'
  },
  input: {
    ...formComponent,
    textIndent: length.small
  },
  button: {
    ...formComponent,
    backgroundColor: color.blue,
    color: color.white,
  }
}

export default class Login extends Component {
  constructor(props) {
    super(props)
    this.state = {
      username: '',
      password: ''
    }
  }

  handleSubmit = (event) => {
    const { username, password } = this.state;
    this.props.loginSubmit(username, password)
    event.preventDefault()
  }

  usernameChange = (event) => {
    this.setState({username: event.target.value})
  }

  passwordChange = (event) => {
    this.setState({password: event.target.value})
  }

  render() {
    const { username, password } = this.state
    return (
      <div>
        <form onSubmit={this.handleSubmit} style={styles.form}>
          <input
            type="text"
            value={username}
            placeholder="Username"
            onChange={this.usernameChange}
            style={styles.input}
            className='card'
          />
          <input
            type="password"
            value={password}
            placeholder="Password"
            onChange={this.passwordChange}
            style={styles.input}
            className='card'
          />
        <input type="submit" value="Log In" style={styles.button} className='card'/>
        </form>
      </div>
    )
  }
}
