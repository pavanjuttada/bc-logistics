

var bodyParser = require('body-parser');

var express = require('express');

var app = express();

var logistics = require('./routes/logistics.js');
var path = require('path');



app.use(bodyParser.json());

app.use(bodyParser.urlencoded({extended: false}));
app.set('view engine', 'html');
app.engine('html', require('ejs').renderFile);
app.set('views', __dirname+ '/public');
app.use(express.static(__dirname +'/public'));



app.use('/', logistics);



// catch 404 and forward to error handler

app.use(function (req, res, next) {

	var err = new Error('Not Found');

	err.status = 404;

	next(err);

});



// error handler

app.use(function (err, req, res, next) {

	// set locals, only providing error in development

	res.locals.message = err.message;

	res.locals.error = req.app.get('env') === 'development' ? err : {};



	// render the error page

	res.status(err.status || 500);

	// res.render('error');
	res.json({
		message: err.message,
		error: err
	});

});



//Listening on port 4000

app.listen('4000', function (e, r) {

	console.log("Listening on port 4000");

});



module.exports = app;


