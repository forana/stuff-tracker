angular.module("stuff-tracker")
.config(function($stateProvider) {
    $stateProvider
    .state('recent-adds', {
        url: "/",
        templateUrl: "/components/recent-adds/recent-adds.html",
        controller: controller
    });
    console.log("yo");
});

const controller = function($scope, Items) {
    $scope.items = [];

    Items.getRecent(10)
    .then(function(items) {
        $scope.items = items;
    });
};
