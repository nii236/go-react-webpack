import React, { Component } from 'react';
import Add from 'app/Components/Add';
import Multiply from 'app/Components/Multiply';
import Header from 'app/Components/Header';

export default class App extends Component {
  render() {
    return (
      <div className="react-app">
        <Header/>
        <div>
          <Add/>
          <Multiply/>
        </div>
      </div>
    );
  }
}
