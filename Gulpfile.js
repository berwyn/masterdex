var gulp 	= require('gulp'),
	sass 	= require('gulp-sass'),
	prefix 	= require('gulp-autoprefixer'),
	jshint 	= require('gulp-jshint'),
	stylish = require('jshint-stylish'),
	watch	= require('gulp-watch'),
	env 	= process.env.NODE_ENV || 'development';

var paths = {
	js: ['**/*.js', '!node_modules/**/*.js'],
	css: 'assets/css/*.scss'
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

gulp.task('stylesheets', function() {
	cssFunc(gulp.src(paths.css));
});

gulp.task('jshint', function() {
	jsFunc(gulp.src(paths.js));
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