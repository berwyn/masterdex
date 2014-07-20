module.exports = {
  up: function(migration, DataTypes, done) {
    migration.addIndex(
		'Species',
		['dexNumber'],
		{
			indexName: 'Species_DexNumber',
			indiciesType: 'UNIQUE'
		}
	).complete(done);
  },
  down: function(migration, DataTypes, done) {
    migration.removeIndex('Species_DexNumber').comlete(done);
  }
}
