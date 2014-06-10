module.exports = function(grunt) {
	'use strict';

	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),
		jshint: {
			all: ['Gruntfile.js', 'app.js', 'controllers/**/*.js'],
			options: {
				node: true,
			    curly: true,
			    eqeqeq: true,
			    eqnull: true,
			    browser: true
			}
		},
		watch: {
			jshint: {
				files: ['**/*.js', '!**/node_modules/**'],
				tasks: ['jshint'],
				options: {
					interrupt: true
				}
			}
		}
	});

	require('load-grunt-tasks')(grunt);
};