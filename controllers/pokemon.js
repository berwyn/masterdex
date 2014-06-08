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

};

PokemonController.prototype.create = function create(req, res) {

};

PokemonController.prototype.get = function get(req, res) {

};

PokemonController.prototype.update = function update(req, res) {

};

PokemonController.prototype.remove = function remove(req, res) {

};

module.exports = PokemonController;