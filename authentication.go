package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"

	"math/rand"

	"github.com/goadesign/goa"
	"github.com/nii236/go-react-webpack/app"
	"github.com/nii236/go-react-webpack/secrets"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/GoogleCallback",
		ClientID:     secrets.ClientID,
		ClientSecret: secrets.ClientSecret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}

	facebookOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/facebookCallback",
		ClientID:     secrets.ClientID,
		ClientSecret: secrets.ClientSecret,
		Scopes:       []string{},
		Endpoint:     facebook.Endpoint,
	}

	githubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/githubCallback",
		ClientID:     secrets.ClientID,
		ClientSecret: secrets.ClientSecret,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString = string(rand.Int63())
)

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>
`

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service) *AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("authentication")}
}

// CallbackResponseFromGoogle runs the Callback response from Google action.
func (c *AuthenticationController) CallbackResponseFromGoogle(ctx *app.CallbackResponseFromGoogleAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request

	code := r.FormValue("code")
	state := r.FormValue("state")
	fmt.Fprintf(w, "Hello! Code is: %s\n", code)

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "{Content: %s, AuthToken: %s}", contents, token.AccessToken)
	return nil
}

// LogIntoGoogle runs the Log into Google action.
func (c *AuthenticationController) LogIntoGoogle(ctx *app.LogIntoGoogleAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

// LogIntoFacebook runs the Log into Google action.
func (c *AuthenticationController) LogIntoFacebook(ctx *app.LogIntoFacebookAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request
	url := facebook.Endpoint.AuthURL
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

// CallbackResponseFromFacebook runs the Callback response from Facebook action.
func (c *AuthenticationController) CallbackResponseFromFacebook(ctx *app.CallbackResponseFromFacebookAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request

	code := r.FormValue("code")
	state := r.FormValue("state")
	fmt.Fprintf(w, "Hello! Code is: %s", code)

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}
	return nil
}

// LogIntoGithub runs the Log into Github action.
func (c *AuthenticationController) LogIntoGithub(ctx *app.LogIntoGithubAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request
	url := github.Endpoint.AuthURL
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

// CallbackResponseFromGithub runs the Callback response from Google action.
func (c *AuthenticationController) CallbackResponseFromGithub(ctx *app.CallbackResponseFromGithubAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request

	code := r.FormValue("code")
	state := r.FormValue("state")
	fmt.Fprintf(w, "Hello! Code is: %s", code)

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}
	return nil
}
