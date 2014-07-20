var gulp 	= require('gulp'),
	sass 	= require('gulp-sass'),
	prefix 	= require('gulp-autoprefixer'),
	jshint 	= require('gulp-jshint'),
	stylish = require('jshint-stylish'),
	watch	= require('gulp-watch'),
	_		= require('lodash-node'),
	env 	= process.env.NODE_ENV || 'development';

var paths = {
	js: ['**/*.js', '!node_modules/**/*.js', '!bower_components/**/*.js', '!static/**/*.js'],
	css: 'assets/css/*.scss',
	components: {
		'bootstrap/dist/css/bootstrap.css': './static/css',
		'bootstrap/dist/js/bootstrap.js': './static/js',
		'bootstrap/fonts/*': './static/fonts',
		'jquery/dist/jquery.js': './static/js'
	}
};

var cssFunc = function cssFunc(source) {
	var options = {
		imagePath: 'static/img'
	};
	source
		.pipe(sass(options))
		.pipe(prefix())
		.pipe(gulp.dest('static/css'));
};

var jsFunc = function jsFunc(source) {
	var options = {
		node: true,
		curly: true,
		eqeqeq: true,
		eqnull: true,
		browser: true
	};
	source
		.pipe(jshint(options))
		.pipe(jshint.reporter(stylish));
};

gulp.task('stylesheets', function(cb) {
	cssFunc(gulp.src(paths.css));
	cb();
});

gulp.task('jshint', function(cb) {
	jsFunc(gulp.src(paths.js));
	cb();
});

gulp.task('components', function(cb) {
	_(paths.components).keys().each(function(key) {
		gulp.src('./bower_components/' + key)
			.pipe(gulp.dest(paths.components[key]));
	});
	cb();
});

gulp.task('compile', ['stylesheets', 'jshint', 'components'], function() {
	//noop
});

gulp.task('watch', function() {
	gulp.src(paths.js)
		.pipe(watch(function(files) {
			jsFunc(files);
		}));
	gulp.src(paths.css)
		.pipe(watch(function(files) {
			cssFunc(files);
		}));
});