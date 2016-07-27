'use strict';

import { Platform } from 'react-native';
import io from 'socket.io-client/socket.io';
import { SERVER_URL } from '../../credentials/server-info';

const socket = io.connect(SERVER_URL, {
  jsonp: false,
  transports: ['websocket']
});

socket.on('get info from client', (data) => {
  if (data === 'GET INFO') {
    socket.emit('send info to server', {
      clientMachine: (Platform.OS + " running React Native")
    });
  }
})

export default socket;
