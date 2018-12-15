var express = require('express')
var expressSession = require('express-session')
var http = require('http')
var io = require('socket.io')

var session = expressSession({ /* configuration */ })

// middleware
app.use(session)

// routes
app.get('/', function() {})

// setup servers
var server = http.createServer(app);
var sio = io.listen(server);

// setup socket auth with sessions
sio.set('authorization', function(handshake, accept) {
  session(handshake, {}, function (err) {
    if (err) return accept(err)
    var session = socket.handshake.session;
    // check the session is valid
    accept(null, session.userid != null)
  })
})

// setup socket connections to have the session on them
sio.sockets.on('connection', function (socket) {
  session(socket.handshake, {}, function (err) {
    if (err) { /* handle error */ }
    var session = socket.handshake.session;
    // do stuff

    // alter session
    session.userdata = mydata

    // and save session
    session.save(function (err) { /* handle error */ })
  })
})

