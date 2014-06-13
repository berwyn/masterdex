'use strict';

var express 	= require('express'),
	hbs			= require('hbs'),
	hbsutils	= require('hbs-utils')(hbs),
	less		= require('less-middleware'),
	app			= express(),
	router  	= express.Router(),
	port		= process.env.PORT || 8080;

app.set('port', port);
app.set('view engine', 'hbs');
app.set('views', __dirname + '/views');

app.use(less(__dirname + '/static'));
app.use(express.static(__dirname + '/static'));
app.use(function(req, res, next) {
	app.locals.meta = { path: req.path };
	next();
});

hbs.localsAsTemplateData(app);
hbs.registerHelper('eq', function(first, second, options) {
	return (first === second)? options.fn(this) : options.inverse(this);
});

var partialDir = __dirname + '/views/partials';
if(app.get('env') === 'development') {
	hbsutils.registerWatchedPartials(partialDir);
} else {
	hbsutils.registerPartials(partialDir, {precompile: true});
}

var controllers = require('./controllers')(router);
var db 			= require('./model');

app.use('/', router);
app.listen(port, function() {
	console.log('Masterdex booting in [' + app.get('env') + '] on :' + port);
});