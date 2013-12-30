var masterdex = angular.module('masterdex', ['ngRoute']);

masterdex.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
  $locationProvider.html5Mode(true);
  $routeProvider
    .when('/', {
      templateUrl: 'root.html',
      controller: 'RootCtrl'
    })
    .when('/pokemon', {
      templateUrl: 'pokemon.html',
      controller: 'PkmnCtrl'
    })
    .when('/item', {
      templateUrl: 'item.html',
      controller: 'ItemCtrl'
    })
    .when('/about', {
      templateUrl: 'about.html',
      controller: 'AboutCtrl'
    });
}]);

masterdex.controller('AppCtrl', ['$scope', '$location', function($scope, $location) {
  $scope.$watch(function() {
    return $location.path();
  }, function(path) {
    if(path === '/') {
      $('body').addClass('home');
    } else {
      $('body').removeClass('home');
    }
  });
}]);

masterdex.controller('RootCtrl', ['$scope', function($scope) {

}]);

masterdex.controller('PkmnCtrl', ['$scope', function($scope) {

}]);

masterdex.controller('ItemCtrl', ['$scope', function($scope) {

}]);

masterdex.controller('AboutControl', ['$scope', function($scope) {

}]);