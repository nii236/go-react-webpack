import React, { Component } from 'react';
import axios from 'axios';

export default class Operands extends Component {
  constructor(props) {
    super(props);
    this.state = {
      left: 0,
      right: 0,
      answer: 'Something'
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit() {
    console.log('Clicked');
    const left = this.state.left;
    const right = this.state.right;
    axios.get('http://localhost:8080/add/' + left + '/' + right)
      .then((response) => {
        console.log(response);
        this.setState({answer: response.data})
      })
      .catch( (response) => {
        console.log("DERP");
        this.setState({answer: "API Error"})
      });

  }

  handleChange(position, event) {
    if (position === 'left') {
      this.setState({left: event.target.value})
    } else if (position === 'right') {
      this.setState({right: event.target.value})
    } else {
      console.error('Invalid position');
    }
  }

  render() {
    return (
      <div>
        Please enter your operands
        <br/>
        <input className='btn btn-default' onChange={this.handleChange.bind(this, 'left')}/>
        <br/>
        +
        <br/>
        <input className='btn btn-default' onChange={this.handleChange.bind(this, 'right')}/>
        <br/>
        <div className='h1'>
          = {this.state.answer}
        </div>
        <br/>
        <button className='btn btn-default' onClick={this.handleSubmit}>
          Submit
        </button>
      </div>
    );
  }
}
