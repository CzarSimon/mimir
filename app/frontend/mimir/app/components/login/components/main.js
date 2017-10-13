import React, { Component } from 'react';
import { View, Text, Image, StyleSheet } from 'react-native';
import { SocialIcon } from 'react-native-elements';
import { color, font, length } from '../../../styles/styles';
import logo from '../../../assets/images/mimir-icon.png';

export default class Login extends Component {
  render() {
    const { loginWith } = this.props;
    return (
      <View style={style.container}>
        <Image style={style.logo} source={logo}/>
        <Text style={style.text}>Sign in to mimir</Text>
        <SocialIcon
          type='twitter'
          title='Sign in with twitter'
          fontFamily={font.type.sans.normal}
          fontStyle={{fontSize: font.h5}}
          onPress={() => loginWith('twitter')}
          style={style.button}
          iconSize={30}
          button
        />
        <SocialIcon
          type='facebook'
          title='Sign in with facebook'
          fontFamily={font.type.sans.normal}
          fontStyle={{fontSize: font.h5}}
          onPress={() => loginWith('facebook')}
          style={style.button}
          iconSize={30}
          button
        />
      </View>
    )
  }
}

const style = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    marginHorizontal: length.medium
  },
  logo: {
    marginTop: 2 * length.button,
    marginBottom: length.large
  },
  text: {
    fontFamily: font.type.sans.bold,
    marginBottom: length.large,
    fontSize: font.h1,
    color: color.blue
  },
  button: {
    alignSelf: 'stretch',
    borderRadius: length.mini
  }
})
