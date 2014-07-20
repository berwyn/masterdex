var fs 			= require('fs'),
	path 		= require('path'),
	Sequelize 	= require('sequelize'),
	_ 			= require('lodash'),
	config		= require('../config/config')
	db 			= {};

module.exports = function connect() {
	'use strict';


	var env 		= config[process.env.ENV_VARIABLE || 'development'],
		sequelize 	= new Sequelize(env['database'], env['username'], env['password'], {
			host: env['host'],
			dialect: env['dialect']
		});


	fs
		.readdirSync(__dirname)
		.filter(function(file) {
			return (file.indexOf('.') !== 0) && (file !== 'index.js');
		})
		.forEach(function(file) {
			var model = sequelize.import(path.join(__dirname, file));
			db[model.name] = model;
		});

	Object.keys(db).forEach(function(modelName) {
		if('associate' in db[modelName]) {
			db[modelName].associate(db);
		}
	});

	return _.extend({
		sequelize: sequelize,
		Sequelize: Sequelize
	}, db);
};