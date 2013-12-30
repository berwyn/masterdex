window.onload = function() {
  pokeball = document.getElementById('pokeball');
  if(window.location.pathname === "/") {
    $('body').addClass('home');
  }
};

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

masterdex.controller('RootCtrl', ['$scope', function($scope) {

}]);

masterdex.controller('PkmnCtrl', ['$scope', function($scope) {

}]);

masterdex.controller('ItemCtrl', ['$scope', function($scope) {

}]);

masterdex.controller('AboutControl', ['$scope', function($scope) {

}]);