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
        <Login/>
        <div>
          <i className="fa fa-google-plus-square fa-6x"></i>
          <i className="fa fa-camera-retro fa-lg"></i> fa-lg
          <i className="fa fa-camera-retro fa-2x"></i> fa-2x
          <i className="fa fa-camera-retro fa-3x"></i> fa-3x
          <i className="fa fa-camera-retro fa-4x"></i> fa-4x
          <i className="fa fa-camera-retro fa-5x"></i> fa-5x
          <Add/>
          <Multiply/>
        </div>
      </div>
    );
  }
}
