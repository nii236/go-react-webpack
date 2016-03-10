import React, { Component } from 'react';
import Add from 'app/Components/Add';
import Multiply from 'app/Components/Multiply';
import Header from 'app/Components/Header';
import Login from 'app/Components/Login';
import Auth from 'app/Services/Auth';

export default class App extends Component {
  constructor(props) {
    super(props)
    this.handleClick = this.handleClick.bind(this);
  }
  handleClick() {
    console.log('Attempt login');
    Auth.login();
  }
  render() {
    return (
      <div className="react-app">
        <button onClick={this.handleClick}>hello</button>
        Click <a href="http://localhost:8080/GoogleLogin">here to login</a>
        <Login/>
        <div>
          <Add/>
          <Multiply/>
        </div>
      </div>
    );
  }
}
