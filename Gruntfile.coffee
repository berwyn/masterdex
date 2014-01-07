module.exports = (grunt) ->
  require('load-grunt-tasks')(grunt)

  grunt.initConfig
    pkg: grunt.file.readJSON 'package.json'

    uglify:
      options:
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'

    coffee:
      options:
        join: true
      files:
        'public/js/masterdex.js': ['assets/scripts/masterdex.coffee']

    sass:
      files:
        'public/css/masterdex.css': 'assets/stylesheets/masterdex.sass'