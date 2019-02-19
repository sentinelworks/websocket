var express = require('express');
var tokenMiddleware = require('../middleware/token');

//var User = require('../models/user');
//var Event = require('../models/event');

var router = express.Router();

router.get('/*', function getEvent(request, response) {
    console.log(request.params);
    console.log("URL: %s", request.originalUrl);
    console.log(request.body);
    response.json({
      success: true,
      event:"getworks" 
    });

/*

  Event
  .findOne({
    id: request.params.eventId
  })
  .populate('tickets')
  .exec(function handleQuery(error, event) {

    if (error) {
      response.status(500).json({
        success: false,
        message: 'Internal server error'
      });

      throw error;
    }

    if (! event) {
      response.status(404).json({
        success: false,
        message: "Can't find event with id " + request.params.eventId + "."
      });

      return;
    }

    response.json({
      success: true,
      event: event
    });
  });
*/

});

router.post('/tokens', function (req, res) {
    var post_body = req.body;
    console.log("URL: %s", req.originalUrl);
    console.log(post_body);
    res.writeHead(200, {'Content-Type': 'application/json'});
    var response = {"data": {"creation_time": "1487002526", "last_modified": "1487002526", "app_name": "", "session_token": "b8b2b634ed272bc74eb675a71bd3a41b", "id": "19178950a90fb3359a00000000000000000000059a", "username": "admin", "source_ip": "10.18.183.184"}};
    res.end(JSON.stringify(response));
})


router.post('/*', tokenMiddleware.verifyToken, function createEvent(request, response) {

    console.log("URL: %s", request.originalUrl);
    console.log(request.body);
    //var obj = JSON.parse(request.body);
    response.json({
          success: true,
          event: "post works" 
        });

/*
  // find the user
  User.findOne({
    username: request.body.username
  }, function handleQuery(error, user) {

    if (error) {
      response.status(500).json({
        success: false,
        message: 'Internal server error'
      });

      throw error;
    }

    if (! user) {
      response.status(404).json({
        success: false,
        message: "Can't find user with username " + request.body.username + "."
      });

      return;
    }

    console.log(user);

    var event = new Event({
      description: request.body.description,
      price: request.body.price,
      createdBy: user._id,
      lastName: request.body.lastName
    });

    event.save(function (error) {

      if (error) {
        response.status(500).json({
          success: false,
          message: 'Internal server error'
        });

        throw error;
      }

      user.events.push(event);

      user.save(function (error) {

        if (error) {
          response.status(500).json({
            success: false,
            message: 'Internal server error'
          });

          throw error;
        }

        response.json({
          success: true,
          event: event
        });
      });
    });
  });
*/

});

module.exports = router;
