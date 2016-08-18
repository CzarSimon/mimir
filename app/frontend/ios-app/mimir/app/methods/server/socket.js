'use strict';

import { Platform } from 'react-native';
//The line below needs to be commented out before running with debugger.
//window.navigator.userAgent = "react-native";
let io = require('socket.io-client/socket.io');
import { SERVER_URL } from '../../credentials/server-info';

const socket = io.connect(SERVER_URL, {
  jsonp: false,
  transports: ['websocket']
});

socket.on("GET_CLIENT_INFO", (data) => {
  if (data === 'GET INFO') {
    socket.emit("DISPATCH_CLIENT_INFO", {
      client_machine: (Platform.OS + " running React Native")
    });
  }
})

export default socket;
