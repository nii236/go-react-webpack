package client

import (
	"io"
	"net/http"
	"net/url"
)

// login lets the user login to their previously registered account
func (c *Client) LoginAuthentication(path string, password string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	values := u.Query()
	values.Set("password", password)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// logout lets the user logout to their previously registered account
func (c *Client) LogoutAuthentication(path string, password string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	values := u.Query()
	values.Set("password", password)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// signup lets a user create a new account
func (c *Client) SignupAuthentication(path string, password string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	values := u.Query()
	values.Set("password", password)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
