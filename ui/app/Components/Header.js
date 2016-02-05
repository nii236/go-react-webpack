import React, { Component } from 'react';
import Add from 'app/Components/Add';
import Multiply from 'app/Components/Multiply';


export default class Header extends Component {
  render() {
    return (
      <div className="Header">
        <ul>Calculator app</ul>
        <ul>Sign up</ul>
        <ul>Log in</ul>
      </div>
    );
  }
}
