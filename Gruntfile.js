module.exports = function(grunt) {
	'use strict';

	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),
		
		jshint: {
			all: [
				'**/*.js', 
				'!**/node_modules/**', 
				'!**/vendor/**', 
				'!**/static/**'
			],
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
				files: [
					'**/*.js', 
					'!**/node_modules/**', 
					'!**/vendor/**',
					'!**/static/**'
				],
				tasks: ['jshint'],
				options: {
					interrupt: true
				}
			},
			stylesheets: {
				files: [
					'assets/css/**.scss'
				],
				tasks: ['sass:dev'],
				options: {
					interrupt: true
				}
			}
		},
		
		sass: {
			dist: {
				options: {
					imagePath: '/static/img'
				},
				files: {
					'static/css/mane.css': 'assets/css/mane.scss'
				}
			},
			dev: {
				options: {
					imagePath: '/static/img',
					sourceMap: true
				},
				files: {
					'static/css/mane.css': 'assets/css/mane.scss'
				}
			}
		},

		cssmin: {
			dist: {
				files: {
					'static/css/mane.css': 'static/**/*.css'
				}
			}
		},

		autoprefixer: {
			static: {
				expand: true,
				cwd: 'static',
				src: ['**/*.css'],
				dest: 'static'
			}
		},

		copy: {
			vendor: {
				cwd: 'vendor',
				src: ['**'],
				dest: 'static',
				expand: true
			}
		},

		clean: {
			static: {
				src: ['static']
			}
		}
	});

	var env = process.env.NODE_ENV || 'development';
	grunt.registerTask(
		'stylesheets',
		'Compiles the stylesheets',
		function() {
			if(env === 'development') {
				grunt.task.run('sass:dev');
			} else {
				grunt.task.run('sass:dist');
			}
			grunt.task.run('autoprefixer');
			if(env === 'development') {
				grunt.task.run('cssmin');
			}
		}
	);

	grunt.registerTask(
		'build',
		'Compiles front-end assets',
		function() {
			grunt.task.run('clean');
			if(env === 'development') {
				grunt.task.run('stylesheets:dev');
			} else {
				grunt.task.run('stylesheets:dist');
			}
			grunt.task.run('copy');
		}
	);

	require('load-grunt-tasks')(grunt);
};