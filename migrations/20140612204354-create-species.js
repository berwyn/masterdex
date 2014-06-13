module.exports = {
  up: function(migration, DataTypes, done) {
    migration.addTable(
    	'species',
    	{
    		id: {
		      type: DataTypes.INTEGER,
		      primaryKey: true,
		      autoIncrement: true
		    },
		    createdAt: {
		      type: DataTypes.DATE
		    },
		    updatedAt: {
		      type: DataTypes.DATE
		    },
		    name: {
		    	type: DataTypes.STRING
		    },
			description: {
				type: DataTypes.STRING
			},
			dexNumber: {
				type: DataTypes.INTEGER
			}
		 },
		 {

		 }
	);
	migration.addIndex(
		'species',
		['dex_number'],
		{
			indexName: 'Species_DexNumber',
			indiciesType: 'UNIQUE'
		}
	);
    done();
  },
  down: function(migration, DataTypes, done) {
  	migration.removeIndex('species', 'Species_DexNumber');
  	migration.dropAllTables();
    done();
  }
};