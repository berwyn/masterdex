var fs 		= require('fs'),
	path 	= require('path');

module.exports = function(router) {
	'use strict';

	fs
		.readdirSync(__dirname)
		.filter(function(file) {
			return (file.indexOf('.') !== 0) && (file !== 'index.js');
		})
		.forEach(function(file) {
			var controller = require(path.join(__dirname, file));
			if('register' in controller) {
				controller.register(router);
			}
		});
};