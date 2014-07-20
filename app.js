'use strict';

var express 	= require('express'),
	hbs			= require('hbs'),
	hbsutils	= require('hbs-utils')(hbs),
	Negotiator	= require('negotiator'),
	bodyParser	= require('body-parser'),
	controllers = require('./controllers'),
	app			= express(),
	router  	= express.Router(),
	port		= process.env.PORT || 8080;

// Configure
app.set('port', port);
app.set('view engine', 'hbs');
app.set('views', __dirname + '/views');

// Middleware
app.use(express.static(__dirname + '/static'));
app.use(bodyParser());
app.use(function(req, res, next) {
	res.locals.meta = { path: req.path };
	next();
});
app.use(router);

// Handlebars helpers
hbs.registerHelper('eq', function(first, second, options) {
	return (first === second)? options.fn(this) : options.inverse(this);
});
hbs.registerHelper('list', function(context, options) {
  var ret = "<ul>";

  for(var i=0, j=context.length; i<j; i++) {
    ret = ret + "<li>" + options.fn(context[i]) + "</li>";
  }

  return ret + "</ul>";
});

// Views
var partialDir = __dirname + '/views/partials';
if(app.get('env') === 'development') {
	hbsutils.registerWatchedPartials(partialDir);
} else {
	hbsutils.registerPartials(partialDir, {precompile: true});
}

// Serve
controllers(router);
app.listen(port, function() {
	console.log('Masterdex booting in [' + app.get('env') + '] on :' + port);
});