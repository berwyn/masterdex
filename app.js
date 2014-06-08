'use strict';

var express 	= require('express'),
	app			= express(),
	router  	= express.Router(),
	port		= process.env.PORT || 8080;

app.set('port', port);

var controller 	= require('./controllers')(router);

app.use('/', router);
app.listen(port);