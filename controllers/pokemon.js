//Varous NodeJS modules
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

/**
 * Dummy data being used until data storage is sorted
 * @type {Array}
 */
var dummy = [
	{
		name: {
			en: 'Bulbasaur',
			jp: 'フシギダネ',
			fr: 'Bulbizarre',
			de: 'Bisasam',
			kr: '이상해씨'
		},
		dexNumber: 1,
		description: 'The seed pokemon'
	}, {
		name: {
			en: 'Ivysaur',
			jp: 'フシギソウ',
			fr: 'Herbizarre',
			de: 'Bisaknosp',
			kr: '이상해풀'
		},
		dexNumber: 2,
		description: 'The seed pokemon'
	}, {
		name: {
			en: 'Venusaur',
			jp: 'フシギバナ',
			fr: 'Florizarre',
			de: 'Bisaflor',
			kr: '이상해꽃',
		},
		dexNumber: 3,
		description: 'The seed pokemon'
	}
];

function PokemonController(){}

/**
 * Registers the controller with the provided Express router
 * @param  {express.Router} 
 *         router the router to register with
 */
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
 * @param  {Object} entity 
 *         The entity who's properties we want to properties we want to take
 * @return {Object}        
 *         An object with publicly-allowable properties
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
 * @param  {express.Request} req      
 *         The incoming request
 * @param  {express.Response} res      
 *         The outbound response
 * @param  {Object} entity   
 *         The raw entity to provide to the template
 * @param  {String} template 
 *         The template to render
 */
PokemonController.prototype.render = function render(req, res, entity, template) {
	var payload,
		responseType = new Negotiator(req).mediaType(mediaTypes);

	if(_.isArray(entity)) {
		payload = { collection: [] };
		entity.forEach(function(el) {
			payload.collection.push(this.buildPayload(el));
		}.bind(this));
	} else {
		payload = { pokemon: this.buildPayload(entity) };
	}

	switch(responseType) {
		case 'application/json':
			res.json(payload);
			break;
		default:
			if(req.method === 'GET') {
				res.render(template, payload);
			} else {
				res.redirect('/pokemon/' + ('00' + entity.dexNumber).slice(-3));
			}
			break;
	}
};

PokemonController.prototype.index = function index(req, res) {
	this.render(req, res, dummy, 'pokemon/index');
};

PokemonController.prototype.create = function create(req, res) {
	res.send('[POST] /pokemon');
};

PokemonController.prototype.get = function get(req, res) {
	this.render(req, res, dummy[0], 'pokemon/show');
};

PokemonController.prototype.update = function update(req, res) {
	res.send('[' + req.method + '] /pokemon/' + req.params.id);
};

PokemonController.prototype.remove = function remove(req, res) {
	res.send('[DELETE] /pokemon/' + req.params.id);
};

module.exports = PokemonController;