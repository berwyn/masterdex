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
      compile:
        options:
          sourceComments: 'map'
          sourceMap: 'public/css'
        cwd: 'assets/stylesheets'
        src: '*.scss'
        dest: 'assets/stylesheets'
        expand: true
        ext: '.css'

    watch:
      scripts:
        files: ['assets/scripts/*.coffee']
        tasks: ['coffee']
      styles:
        files: ['assets/stylesheets/*.sass']
        tasks: ['sass']

  grunt.registerTask 'default', ['coffee', 'sass', 'watch']
  grunt.registerTask 'compile', ['coffee', 'sass']
  grunt.registerTask 'heroku', ['coffee', 'sass']