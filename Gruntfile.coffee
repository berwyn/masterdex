module.exports = (grunt) ->
  require('load-grunt-tasks')(grunt)

  grunt.initConfig
    pkg: grunt.file.readJSON 'package.json'

    uglify:
      options:
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'

    coffee:
      compile:
        options:
          join: true
        files:
          'public/js/masterdex.js': ['assets/scripts/masterdex.coffee']

    sass:
      dist:
        options:
          sourcemap: true
          style: 'compact'
        files:
          'public/css/masterdex.css': 'assets/stylesheets/masterdex.sass'

    watch:
      scripts:
        files: ['assets/scripts/*.coffee']
        tasks: ['coffee']
      styles:
        files: ['assets/stylesheets/*.sass']
        tasks: ['sass']

  grunt.registerTask 'default', ['coffee', 'sass', 'watch']