angular.module("stuff-tracker", ["ui.router", "items"])
.config(function ($urlRouterProvider, $locationProvider, $httpProvider) {
    $urlRouterProvider.otherwise('/');
    $locationProvider.html5Mode({
        enabled: false,
        requireBase: false
    });
    $httpProvider.defaults.headers.common['Content-Type'] = 'application/json';
});
