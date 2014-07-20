'use strict';

var MEDIA_TYPES = ['application/json', 'text/html'];

var express 	= require('express'),
	hbs			= require('hbs'),
	hbsutils	= require('hbs-utils')(hbs),
	less		= require('less-middleware'),
	Negotiator	= require('negotiator'),
	bodyParser	= require('body-parser'),
	app			= express(),
	router  	= express.Router(),
	port		= process.env.PORT || 8080;

app.set('port', port);
app.set('view engine', 'hbs');
app.set('views', __dirname + '/views');

app.use(less(__dirname + '/static'));
app.use(express.static(__dirname + '/static'));
app.use(bodyParser());
app.use(function(req, res, next) {
	app.locals.meta = { path: req.path };
	next();
});

app.use(function(req, res, next) {
	next();
	if(res.locals.entity) {
		var mediaType = new Negotiator(req).mediaType(MEDIA_TYPES);
		switch(mediaType) {
			case 'application/json':
				res.send(JSON.stringify(res.locals.entity));
				break;
			case 'text/html':
			default:
				res.render(res.locals.template);
				break;
		}
	}
})

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