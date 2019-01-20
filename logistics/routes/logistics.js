var express = require('express');

var router = express.Router();

var logisticsQueryController=require('../query.js');
var logisticsInvokeController=require('../invoke.js');


router.post('/api/getShipments', logisticsQueryController.getShipments, function (res, req, next) {
	next();
});

router.post('/api/newShipment', logisticsInvokeController.createOrUpdateShipment, function (res, req, next) {
	next();
});

router.post('/api/logTimeRaster', logisticsInvokeController.createOrUpdateShipment, function (res, req, next) {
	next();
});

router.post('/api/updateShipment', logisticsInvokeController.createOrUpdateShipment, function (res, req, next) {
	next();
});



router.get('/home', function(req, res) {
	 res.render('html/index.html');  // load the single view file (angular will handle the page changes on the front-end)
});
router.get('/login', function(req, res) {
	 res.render('html/login.html');  // load the single view file (angular will handle the page changes on the front-end)
});
/*app.get('*', function(req, res) {
        res.sendfile('./public/index.html'); // load the single view file (angular will handle the page changes on the front-end)
    });*/

module.exports = router;