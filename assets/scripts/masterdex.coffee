masterdex = angular.module 'masterdex', ['ngRoute']

masterdex.config ['$routeProvider', '$locationProvider', ($routeProvider, $locationProvider) ->
  $locationProvider.html5Mode true
  $routeProvider
    .when '/',
      templateUrl: 'root.html'
      controller: 'RootCtrl'
    .when '/pokemon',
      templateUrl: 'pokemon.html'
      controller: 'PkmnCtrl'
    .when '/item',
      templateUrl: 'item.html'
      controller: 'ItemCtrl'
    .when '/about',
      templateUrl: 'about.html'
      $controller: 'AboutCtrl'
]

masterdex.controller 'AppCtrl', ['$scope', '$location', ($scope, $location) ->
  $scope.$watch () ->
    return $location.path()
  , () ->
    if path == '/'
      $('body').addClass 'home'
    else
      $('body').removeClass 'home'
]

masterdex.controller 'RootCtrl', ['$scope', ($scope) ->
]

masterdex.controller 'PkmnCtrl', ['$scope', ($scope) ->
]

masterdex.controller 'ItemCtrl', ['$scope', ($scope) ->
]

masterdex.controller 'AboutCtrl', ['$scope', ($scope) ->
]