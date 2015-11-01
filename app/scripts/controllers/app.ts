/// <reference path="../app.ts" />

'use strict';

module itcApp {
  export interface IAppScope extends ng.IScope {
    awesomeThings: any[];
  }

  export class AppCtrl {

    constructor (private $scope: IAppScope) {
      $scope.awesomeThings = [
        'HTML5 Boilerplate',
        'AngularJS',
        'Karma'
      ];
    }
  }
}

angular.module('itcApp')
  .controller('AppCtrl', itcApp.AppCtrl);
