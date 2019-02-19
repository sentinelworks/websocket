var express = require('express');
var router = express.Router();
var store = require('../store/db');
var redis = require("redis");

router.get('/*', function (request, response) {
    console.log(request.params);
    console.log("GUI URL: %s", request.originalUrl);
    console.log(request.body);
    response.json({
        success: true,
        event:"getworks" 
    });
});

router.post('/*', function (request, response) {
    console.log("URL: %s", request.originalUrl);
    console.log(request.body);
    store.conn.send(JSON.stringify(request.body));
    console.log("waiting here");

    var sub = redis.createClient();
    sub.on("message", function (channel, message) {
        console.log("Message " + channel + ": " + message);
        response.json(message);
        sub.unsubscribe();
        sub.quit();
    });
    sub.subscribe("infosight", 1);
});

module.exports = router;

