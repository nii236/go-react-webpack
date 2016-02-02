import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import axios from 'axios';

ReactDOM.render(<App />, document.getElementById('root'));

axios.get('http://localhost:8080/add/1/2')
  .then(function(response) {
    console.log(response);
  })
  .catch(function (response) {
    console.log(response);
  });
