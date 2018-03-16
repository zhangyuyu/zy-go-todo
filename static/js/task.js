/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function TaskCtrl($scope, $http) {
    $scope.tasks = [];
    $scope.working = false;

    var logError = function (data, status) {
        console.log('code ' + status + ': ' + data);
        $scope.working = false;
    };

    var refresh = function () {
        return $http.get('/tasks').success(function (data) {
            $scope.tasks = data;
        }).error(logError);
    };

    $scope.addTask = function () {
        $scope.working = true;
        $http.post('/tasks/', {Title: $scope.todoText}).error(logError).success(function () {
            refresh().then(function () {
                $scope.working = false;
                $scope.todoText = '';
            })
        });
    };

    $scope.toggleDone = function (task) {
        data = {Id: task.Id, Title: task.Title, Done: !task.Done}
        $http.put('/tasks/' + task.Id, data).error(logError).success(function () {
            task.Done = !task.Done
        });
    };

    $scope.deleteTask = function (task) {
        $http.delete('/tasks/' + task.Id).error(logError).success(function () {
            refresh().then(function () {
                $scope.working = false;
                $scope.todoText = '';
            })
        });
    };

    $scope.logout = function () {
        $http.get('/logout').error(logError).success(function () {
            console.log("============logout")
        });
    };

    refresh().then(function () {
        $scope.working = false;
    });
}