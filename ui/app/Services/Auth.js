import request from 'reqwest';
import when from 'when';
import {LOGIN_URL, SIGNUP_URL} from '../constants/Login';
import axios from 'axios';

class AuthService {
  // Payload needs to contain:
  //   scope
  //   state
  //   redirect_uri
  //   response_type: token
  //   client_id
  //   nonce
  login(url, payload) {
    axios.get('https://accounts.google.com/o/oauth2/v2/auth',
      {
        scope: ['email'],
        state: ['Hallo callback'],
        redirect_uri: 'http://localhost:8080/GoogleCallback',
        response_type: 'token',
        client_id: '824271538099-i27kba1jtso706f6e2t92u4qtn8kr2fd.apps.googleusercontent.com',
        nonce: 1
      })
      .then(function(response) {
        console.log(response);
      })
      .catch(function(response) {
        console.log(response);
      });
  }
}

export default new AuthService()
