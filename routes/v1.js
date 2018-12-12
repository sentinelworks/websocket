var express = require('express');
var router = express.Router();
var store = require('../store/db');
var res;

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
    //var obj = JSON.parse(request.body);
    res = store.conn.remit("message", JSON.stringify(request.body));
    console.log("waiting here");
    res = response;
/*
    response.json({
          success: true,
          event: "post works"
        });
*/
});

module.exports = router;
//module.exports = res;

