module.exports = function(sequelize, DataTypes) {
	var Species = sequelize.define('Species', {
		id: 					DataTypes.INTEGER,
		createdAt: 		DataTypes.DATE,
		updatedAt: 		DataTypes.DATE,
		name: 				DataTypes.STRING,
		description: 	DataTypes.STRING,
		dexNumber: 		DataTypes.INTEGER,
		typeMask: 		DataTypes.INTEGER,
		imageUrl: 		DataTypes.STRING
	}, {

	});
};