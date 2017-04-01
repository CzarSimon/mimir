'use strict';
import { Platform } from 'react-native'
import { DEV_MODE, SERVER_URL } from '../../credentials/config'
let io = require('socket.io-client/socket.io')

if (!DEV_MODE) {
  window.navigator.userAgent = 'react-native'
}

const socket = io.connect(SERVER_URL, {
  jsonp: false,
  transports: ['websocket']
});

socket.on('GET_CLIENT_INFO', (data) => {
  if (data === 'GET INFO') {
    socket.emit('DISPATCH_CLIENT_INFO', {
      client_machine: (Platform.OS + ' running React Native')
    });
  }
})

export default socket;
