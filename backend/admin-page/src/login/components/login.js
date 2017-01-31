import React, { Component } from 'react';

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
        <form onSubmit={this.handleSubmit}>
          <input type="text" value={username} placeholder="Username" onChange={this.usernameChange} />
          <input type="password" value={password} placeholder="Password" onChange={this.passwordChange} />
          <input type="submit" value="Submit" />
        </form>
      </div>
    )
  }
}
