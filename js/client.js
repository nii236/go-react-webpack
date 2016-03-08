// This module exports functions that give access to the adder API hosted at localhost:8080.
// It uses the axios javascript library for making the actual HTTP requests.
define(['axios'] , function (axios) {
  function merge(obj1, obj2) {
    var obj3 = {};
    for (var attrname in obj1) { obj3[attrname] = obj1[attrname]; }
    for (var attrname in obj2) { obj3[attrname] = obj2[attrname]; }
    return obj3;
  }

  return function (scheme, host, timeout) {
    scheme = scheme || 'http';
    host = host || 'localhost:8080';
    timeout = timeout || 20000;

    // Client is the object returned by this module.
    var client = axios;

    // URL prefix for all API requests.
    var urlPrefix = scheme + '://' + host;

  // add returns the sum of the left and right parameters in the response body
  // path is the request path, the format is "/add/:left/:right"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.addOperands = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // login lets the user login to their previously registered account
  // path is the request path, the format is "/login/:user/:pass"
  // password is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.loginAuthentication = function (path, password, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      params: {
        password: password
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // logout lets the user logout to their previously registered account
  // path is the request path, the format is "/logout/:user/:pass"
  // password is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.logoutAuthentication = function (path, password, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      params: {
        password: password
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // multiply returns the multiplication of the left and right parameters in the response body
  // path is the request path, the format is "/multiply/:left/:right"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.multiplyOperands = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // signup lets a user create a new account
  // path is the request path, the format is "/signup/:user/:pass"
  // password is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.signupAuthentication = function (path, password, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      params: {
        password: password
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }
  return client;
  };
});
