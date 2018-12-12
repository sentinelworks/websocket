var app = require('express')(),
server = require('http').createServer(app),
io = require('Socket.IO').listen(3033);//for mobile
console.log('server is live on '+2700);
server.listen(2700); //for web
