import React, { Component } from 'react';
import Auth from '../Services/Auth'

export default class Login extends React.Component {

  constructor() {
    super()
    this.state = {
      user: '',
      password: ''
    };
  }

  login(e) {
    e.preventDefault();
    Auth.login(this.state.user, this.state.password)
      .catch(function(err) {
        console.error('There was an error logging in');
        console.log('Error logging in', err);
      });
  }

  handleChange(type, event) {
    if (type === 'user') {
      this.setState({
        user: event.target.value
      })
    } else if (type === 'password') {
      this.setState({
        password: event.target.value
      })
    }
  }

  render() {
    return (
      <div className='login jumbotron center-block'>
        <h1>Login</h1>
        <form role='form'>
        <div className='form-group'>
          <label htmlFor='username'>Username</label>
          <input type='text' onChange={this.handleChange.bind(this, 'user')} className='form-control' id='username' placeholder='Username' />
        </div>
        <div className='form-group'>
          <label htmlFor='password'>Password</label>
          <input type='password' onChange={this.handleChange.bind(this, 'password')} className='form-control' id='password' ref='password' placeholder='Password' />
        </div>
        <button type='submit' className='btn btn-default' onClick={this.login.bind(this)}>Submit</button>
      </form>
    </div>
    );
  }
}
