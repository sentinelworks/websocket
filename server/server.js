"use strict";

process.title = 'gmd-proxy';
var webSocketsServerPort = 8080;
var webSocketServer = require('websocket').server;
var http = require('http');

var express = require('express');
var app = express();

var store = require('./store/db');
var cors = require('cors');
var bodyParser = require('body-parser');

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));
app.use(cors());

var v1Routes = require('./routes/v1');
app.use('/', v1Routes);

var gui = http.createServer(app);

gui.listen(7999, function () {
    var host = gui.address().address;
    var port = gui.address().port;

    console.log((new Date()) + ' GUI listening at http://%s:%s', host, port);
});

var redis = require("redis");
var pub = redis.createClient();

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
    console.log((new Date()) + ' Connection accepted.');
    store.conn = connection;

    // user sent some message
    connection.on('message', function(message) {
        console.log(message);
        console.log('message received: ', message.utf8Data);
        if (message.type === 'utf8') { // accept only text
            // first message sent by GMA is hello message
            // All others are response from GUI/infosight
            var msg = JSON.parse(message.utf8Data);
            console.log('message type received: ',msg["msg_type"]);
            console.log(msg);
            if (msg["msg_type"] == "hello") {
                msg["msg_type"] == "good";
                var response = JSON.parse(JSON.stringify(msg));
                response['msg_type'] == "ack";
                console.log(response);
                connection.sendUTF(JSON.stringify(response));
                console.log((new Date()) + ' Array comes up: '
                    + JSON.stringify(response));
            } else {
                console.log("send to infosight");
                pub.publish("infosight", message.utf8Data);
            }
        } else {
            console.log("ERROR: not a text message");
        }
    });

    // user disconnected
    connection.on('close', function(connection) {
        console.log((new Date()) + " Peer "
            + connection.remoteAddress + " disconnected.");
    });
});
