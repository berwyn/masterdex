var expect 		= require('expect.js');
var httpmock 	= require('node-mocks-http');

describe('Pokemon Controller', function() {
	var PokemonController = require('./../controllers/pokemon'),
		controller;

	beforeEach(function() {
		controller = new PokemonController();
	});

	it('should exist', function() {
		expect(controller).to.be.ok();
	});

	describe('collection queries', function() {
		it('should fetch the collection', function() {
			expect(controller.index).to.be.a('function');
			var request = httpmock.createRequest({
				method: 'GET',
				url: '/pokemon',
				headers: { Accept: 'application/json' }
			});
			var response = httpmock.createResponse();
			controller.index(request, response);

			var data = JSON.parse(response._getData());
			expect(data).to.be.ok();
			expect(data.result).to.be.ok();
			expect(data.result).to.be.an('array');
			expect(data.result).to.not.be.empty();
		});

		it('should create new pokemon', function() {
			expect(controller.create).to.be.a('function');
			var request = httpmock.createRequest({
				method: 'POST',
				url: '/pokemon',
				headers: { Accept: 'application/json' }
			})
			var response = httpmock.createResponse();
			controller.create(request, response);

			var data = JSON.parse(response._getData());
			expect(data).to.be.ok();
			expect(data.result).to.be.ok();
			expect(data.result).to.be.an('object');
			// TODO: Compare to our fixture data
		});
	});

	describe('entity queries', function() {
		it('should fetch an entity', function() {
			expect(controller.get).to.be.a('function');
			var request = httpmock.createRequest({
				method: 'GET',
				url: '/pokemon/001',
				headers: { Accept: 'application/json' }
			});
			var response = httpmock.createResponse();
			controller.get(request, response);

			var data = JSON.parse(response._getData());
			expect(data).to.be.ok();
			expect(data.result).to.be.ok();
			expect(data.result).to.be.an('object');
			// TODO: Compare to our fixture data
		});

		it('should update an entity', function() {
			expect(controller.update).to.be.a('function');
			var request = httpmock.createRequest({
				method: 'PATCH',
				url: '/pokemon/001',
				headers: { Accept: 'application/json' }
			});
			var response = httpmock.createResponse();
			controller.update(request, response);

			var data = JSON.parse(response._getData());
			expect(data).to.be.ok();
			expect(data.result).to.be.ok();
			expect(data.result).to.be.an('object');
			// TODO: Compare to our fixture data
		});

		it('should replace an entity', function() {
			expect(controller.UPDATE).to.be.a('function');
			var request = httpmock.createRequest({
				method: 'PUT',
				url: '/pokemon/001',
				headers: { Accept: 'application/json' }
			});
			var response = httpmock.createResponse();
			controller.UPDATE(request, response);

			var data = JSON.parse(response._getData());
			expect(data).to.be.ok();
			expect(data.result).to.be.ok();
			expect(data.result).to.be.an('object');
			// TODO: Compare to our fixture data
		});

		it('should delete an entity', function() {
			expect(controller.remove).to.be.a('function');
			var request = httpmock.createRequest({
				method: 'DELETE',
				url: '/pokemon/001',
				headers: { Accept: 'application/json' }
			});
			var response = httpmock.createResponse();
			controller.remove(request, response);

			var data = JSON.parse(response._getData());
			expect(data).to.be.ok();
			expect(data.result).to.be.ok();
			expect(data.result).to.be.an('object');
			// TODO: Compare to our fixture data
		});
	});
});