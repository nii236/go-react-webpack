import React, { Component } from 'react';
import Add from 'app/Components/Add';
import Multiply from 'app/Components/Multiply';


export default class App extends Component {
  render() {
    return (
      <div className="jumbotron">
        <h1>Adder App!</h1>
        <Add/>
        <Multiply/>
      </div>
    );
  }
}
