'use strict';

var express 	= require('express'),
	ejs			= require('ejs'),
	Negotiator	= require('negotiator'),
	bodyParser	= require('body-parser'),
	controllers = require('./controllers'),
	app			= express(),
	router  	= express.Router(),
	port		= process.env.PORT || 8080;

// Configure
app.set('port', port);
app.set('view engine', 'ejs');
app.set('views', __dirname + '/views');

// Middleware
app.use(express.static(__dirname + '/static'));
app.use(bodyParser());
app.use(function(req, res, next) {
	res.locals.meta = { path: req.path };
	next();
});
app.use(router);

// Serve
controllers(router);
app.listen(port, function() {
	console.log('Masterdex booting in [' + app.get('env') + '] on :' + port);
});