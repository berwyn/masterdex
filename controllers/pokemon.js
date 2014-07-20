function PokemonController(){}

PokemonController.prototype.register = function register(router) {
	router.get('/pokemon', this.index);
	router.post('/pokemon', this.create);
	router.get('/pokemon/:id', this.get);
	router.patch('/pokemon/:id', this.update);
	router.put('/pokemon/:id', this.update);
	router.delete('/pokemon/:id', this.remove);
};

PokemonController.prototype.index = function index(req, res) {
	res.send('[GET] /pokemon');
};

PokemonController.prototype.create = function create(req, res) {
	res.send('[POST] /pokemon');
};

PokemonController.prototype.get = function get(req, res) {
	res.locals.entity = {
		'name': 'Bulbasaur',
		'dexNumber': 1
	}
	res.locals.template = 'pokemon/show';
};

PokemonController.prototype.update = function update(req, res) {
	res.send('[' + req.method + '] /pokemon/' + req.params.id);
};

PokemonController.prototype.remove = function remove(req, res) {
	res.send('[DELETE] /pokemon/' + req.params.id);
};

module.exports = PokemonController;