import React, { Component } from 'react';
import Add from 'app/Components/Add';
import Multiply from 'app/Components/Multiply';
import Header from 'app/Components/Header';
import Login from 'app/Components/Login';

export default class App extends Component {
  render() {
    return (
      <div className="react-app">
        <Login/>
        <div>
          <Add/>
          <Multiply/>
        </div>
      </div>
    );
  }
}
