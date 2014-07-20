var _ 			= require('lodash-node'),
	Negotiator 	= require('negotiator'),
	mediaTypes 	= ['text/html', 'application/json'];

/**
 * An object defining the properties we want to have publicly exposed
 * @type {Object}
 */
var proto = {
	name: '',
	dexNumber: '',
	description: ''
};

function PokemonController(){}

PokemonController.prototype.register = function register(router) {
	router.get(		'/pokemon', 		this.index.bind(this));
	router.post(	'/pokemon', 		this.create.bind(this));
	router.get(		'/pokemon/:id', 	this.get.bind(this));
	router.patch(	'/pokemon/:id', 	this.update.bind(this));
	router.put(		'/pokemon/:id', 	this.update.bind(this));
	router.delete(	'/pokemon/:id', 	this.remove.bind(this));
};

/**
 * Provides a convenience method to create an object with public properties
 * from any given object
 * @param  {Object} entity The entity who's properties we want to properties we want to take
 * @return {Object}        An object with publicly-allowable properties
 */
PokemonController.prototype.buildPayload = function buildPayload(entity) {
	var payload = _(proto).clone();
	_(payload).keys().each(function(key) {
		payload[key] = entity[key];
	});
	return payload;
};

/**
 * Common rendering code
 * @param  {[type]} req      [description]
 * @param  {[type]} res      [description]
 * @param  {[type]} entity   The raw entity to provide to the template
 * @param  {[type]} template The template to render
 * @return {[type]}          [description]
 */
PokemonController.prototype.render = function render(req, res, entity, template) {
	var payload,
		responseType = new Negotiator(req).mediaType(mediaTypes);

	if(_.isArray(entity)) {
		payload = [];
		entity.forEach(function(el) {
			payload.push(this.buildPayload(el));
		}.bind(this));
	} else {
		payload = this.buildPayload(entity);
	}

	switch(responseType) {
		case 'application/json':
			res.json(payload);
			break;
		default:
			if(req.method === 'GET') {
				res.render(template, { entity: payload });
			} else {
				res.redirect('/pokemon/' + ('00' + entity.dexNumber).slice(-3));
			}
			break;
	}
};

PokemonController.prototype.index = function index(req, res) {
	var entities = [{
			name: 'Bulbasaur',
			dexNumber: 1,
			description: 'The leaf pokemon'
		}, {
			name: 'Ivysaur',
			dexNumber: 2,
			description: 'The leaf pokemon'
		}, {
			name: 'Venusaur',
			dexNumber: 3,
			description: 'The leaf pokemon'
		}];
	this.render(req, res, entities, 'pokemon/index');
};

PokemonController.prototype.create = function create(req, res) {
	res.send('[POST] /pokemon');
};

PokemonController.prototype.get = function get(req, res) {
	this.render(req, res, {
		name: 'Bulbasaur',
		dexNumber: 1,
		description: 'The leaf pokemon'
	}, 'pokemon/show');
};

PokemonController.prototype.update = function update(req, res) {
	res.send('[' + req.method + '] /pokemon/' + req.params.id);
};

PokemonController.prototype.remove = function remove(req, res) {
	res.send('[DELETE] /pokemon/' + req.params.id);
};

module.exports = PokemonController;