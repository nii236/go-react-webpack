import React, { Component } from 'react';
import Operands from 'app/Components/Operands';

export default class App extends Component {
  render() {
    return (
      <div className="jumbotron">
        <h1>Adder App!</h1>
        <Operands/>
      </div>
    );
  }
}
