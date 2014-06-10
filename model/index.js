var fs 			= require('fs'),
	path 		= require('path'),
	Sequelize 	= require('sequelize'),
	_ 			= require('lodash'),
	db 			= {};

module.exports = function connect(username, password) {
	var sequelize = new Sequelize(username, password, masterdex);

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
}