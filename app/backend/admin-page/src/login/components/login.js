import React, { Component } from 'react';
import PageTitle from '../../components/util/page-title';
import { length, color } from '../../styles/styles';
import { portraitMode } from '../../methods/helper-methods';

const formComponent = {
  width: '100%',
  height: '6vh',
  marginBottom: length.small,
  fontSize: '1.1em',
  borderStyle: 'none',
}

const styles = {
  form: {
    width: (!portraitMode()) ? '25vw' : '50vw',
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
  },
  title: {
    color: color.blue,
    textAlign: 'center'
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
          <PageTitle customStyle={styles.title} title={"mimir admin page"} />
          <input
            type="text"
            value={username}
            placeholder="Username"
            onChange={this.usernameChange}
            style={styles.input}
            className='card'
            autoFocus
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
