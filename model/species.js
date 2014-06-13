module.exports = function(sequelize, DataTypes) {
	var Species = sequelize.define('Species', {
		name: DataTypes.STRING,
		description: DataTypes.STRING,
		dexNumber: DataTypes.INTEGER
	}, {

	});
};