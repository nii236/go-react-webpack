import React, { Component } from 'react';
import Auth from '../Services/Auth'

export default class Login extends React.Component {
  render() {
    return(
    <div className="navbar navbar-default">
      <a href="http://localhost:8080/FacebookLogin"><i className="fa fa-facebook-square"></i></a>
      <a href="http://localhost:8080/GithubLogin"><i className="fa fa-github-square"></i></a>
      <a href="http://localhost:8080/GoogleLogin"><i className="fa fa-google-plus-square"></i></a>
    </div>)
  }
}
