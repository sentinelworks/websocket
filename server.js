"use strict";

process.title = 'node-chat';
var webSocketsServerPort = 8080;
var webSocketServer = require('websocket').server;
var http = require('http');

var express = require('express');
var cors = require('cors');
var bodyParser = require('body-parser');
var store = require('./store/db');

var app = express();
var socketio = require("socket.io");


app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
  extended: false
}));
app.use(cors());

var v1Routes = require('./routes/v1');
app.use('/', v1Routes);

var gui = http.createServer(app);
var io = socketio(gui);
io.on("connection", socket => {
    console.log("New client connected");
    store.infosight = socket;
});


gui.listen(7999, function () {
    var host = gui.address().address;
    var port = gui.address().port;

    console.log((new Date()) + ' GUI listening at http://%s:%s', host, port);
});



/**
 * Global variables
 */
// latest 100 messages
var history = [ ];
// list of currently connected clients (users)
//var clients = [ ];

/**
 * HTTP server
 */
var server = http.createServer(function(request, response) {
  // Not important for us. We're writing WebSocket server,
  // not HTTP server
});


server.listen(webSocketsServerPort, function() {
  console.log((new Date()) + " Server is listening on port "
      + webSocketsServerPort);
});



/**
 * WebSocket server
 */
var wsServer = new webSocketServer({
  // WebSocket server is tied to a HTTP server. WebSocket
  // request is just an enhanced HTTP request. For more info 
  // http://tools.ietf.org/html/rfc6455#page-6
  httpServer: server
});

// This callback function is called every time someone
// tries to connect to the WebSocket server
wsServer.on('request', function(request) {
    //console.log(request);
  console.log((new Date()) + ' Connection from origin '
      + request.origin + '.');
  // accept connection - you should check 'request.origin' to
  // make sure that client is connecting from your website
  // (http://en.wikipedia.org/wiki/Same_origin_policy)
  var connection = request.accept(null, request.origin); 
  // we need to know client index to remove them on 'close' event
  //var index = clients.push(connection) - 1;
  var userName = false;
  var userColor = false;
  console.log((new Date()) + ' Connection accepted.');
    store.conn = connection;

  // user sent some message
  connection.on('message', function(message) {
    console.log(message);
    console.log('message received: ', message.utf8Data);
    if (message.type === 'utf8') { // accept only text
    // first message sent by user is their name
        // get random color and send it back to the user
        //response.json = { "type" :'color', "owner": "victor"};
        var msg = JSON.parse(message.utf8Data);
        console.log(msg);
        console.log('message type received: ',msg["msg_type"]);
        if (msg["msg_type"] == "hello") {
            msg["msg_type"] == "good";
            var response = JSON.parse(JSON.stringify(msg));
            response['msg_type'] == "ack";
            console.log(response);
            connection.sendUTF(
                JSON.stringify(response));

            console.log((new Date()) + ' User is known as: ' + JSON.stringify(response));
        } else {
            console.log("send to infosight");
            store.infosight.emit("new message", msg);
        }
    }
  });

  // user disconnected
  connection.on('close', function(connection) {
    if (userName !== false && userColor !== false) {
      console.log((new Date()) + " Peer "
          + connection.remoteAddress + " disconnected.");
      // remove user from the list of connected clients
      clients.splice(index, 1);
      // push back user's color to be reused by another user
      colors.push(userColor);
    }
  });
});
