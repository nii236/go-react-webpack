import React, { Component } from 'react';
import Operands from 'app/Operands';

export default class App extends Component {
  render() {
    return (
      <div className="jumbotron">
        <h1>Welcome to Nii's Adder!</h1>
        <Operands/>
      </div>
    );
  }
}
