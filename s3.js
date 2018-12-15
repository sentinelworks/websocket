var app = require('express').createServer()
  , io = require('socket.io').listen(app);

app.listen(3000);

app.get('/webclient', function (req, res) {
    console.log('/webclient');
    //res.sendfile(__dirname + '/web.html');
});

app.get('/mobile', function (req, res) {
    console.log('/mobile');
});

io.sockets.on('connection', function (socket) {
//      socket.emit('pop', { hello: 'world' });
        console.log(', emitting a pop');
    socket.on('push', function (data) {
        console.log('push received, emitting a pop');
        socket.emit('pop', { hello: 'world' });
    });
  socket.on('CH01', function (from, msg) {
    console.log('MSG', from, ' saying ', msg);
  });

});

