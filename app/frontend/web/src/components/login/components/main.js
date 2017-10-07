import React, { Component } from 'react';
import logo from '../../../assets/images/large-icon.svg';
import IoSocialTwitter from 'react-icons/lib/io/social-twitter';
import IoSocialFacebook from 'react-icons/lib/io/social-facebook';
import { portraitMode } from '../../../methods/helper-methods';
import { length, color, font } from '../../../styles/main';

const style = {
  container: {
    display: 'flex',
    flexDirection: 'column'
  },
  header: {
    width: '100%',
    fontSize: font.size.large,
    fontWeight: 'bold',
    textAlign: 'center',
    color: color.blue,
    marginBottom: length.icons.large
  },
  logo: {
    marginTop: length.button,
    height: '150px'
  },
  button: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: length.small,
    marginBottom: length.medium
  },
  icon: {
    flex: 1,
    color: color.white,
    marginTop: '0',
    marginBottom: '0'
  },
  buttonText: {
    flex: 4,
    color: color.white,
    marginTop: '0',
    marginBottom: '0',
    fontSize: font.size.small
  }
}

export default class Login extends Component {
  twitterLogin = () => {
    const { auth } = this.props;
    auth.twitterLogin();
  }

  facebookLogin = () => {
    const { auth } = this.props;
    auth.facebookLogin();
  }

  render() {
    const width = (!portraitMode()) ? '56%' : '90%';
    const marginLeft = (!portraitMode()) ? '27%' : '5%';
    return (
      <div style={{...style.container, width, marginLeft}}>
        <img
          src={logo}
          style={style.logo}
          alt="mimir logo"
          onClick={this.goToHome} />
        <p style={style.header}>Sign in to mimir</p>
        <div
          className="card"
          onClick={this.twitterLogin}
          style={{...style.button, backgroundColor: color.social.twitter}}>
          <IoSocialTwitter size={32} style={style.icon}/>
          <p style={style.buttonText}>Sign in with twitter</p>
        </div>
        <div
          className="card"
          onClick={this.facebookLogin}
          style={{...style.button, backgroundColor: color.social.facebook}}>
          <IoSocialFacebook size={35} style={style.icon}/>
          <p style={style.buttonText}>Sign in with facebook</p>
        </div>
      </div>
    )
  }
}
