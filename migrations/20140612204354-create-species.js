module.exports = {
	up: function(migration, DataTypes, done) {
		migration.createTable(
			'Species',
			{
				name: {
					type: DataTypes.STRING,
					allowNull: false,
					unique: true
				},
				description: {
					type: DataTypes.STRING
				},
				dexNumber: {
					type: DataTypes.INTEGER,
					allowNull: false,
					unique: true
				},
				typeMask: {
					type: DataTypes.INTEGER
				},
				imageUrl: {
					type: DataTypes.STRING
				}
	 		}
		).complete(done);
	},

	down: function(migration, DataTypes, done) {
		migration.dropAllTables().complete(done);
	}
};