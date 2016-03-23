package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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
		ClientID:     secrets.Google.ID,
		ClientSecret: secrets.Google.Secret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}

	facebookEndpoint = &oauth2.Endpoint{
		AuthURL:  "https://www.facebook.com/dialog/oauth",
		TokenURL: "https://graph.facebook.com/v2.3/oauth/access_token",
	}

	facebookOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/facebookCallback",
		ClientID:     secrets.Facebook.ID,
		ClientSecret: secrets.Facebook.Secret,
		Scopes:       []string{},
		Endpoint:     *facebookEndpoint,
	}

	githubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/githubCallback",
		ClientID:     secrets.Github.ID,
		ClientSecret: secrets.Github.Secret,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}

	// Needs to be refactored to change on each request?
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

// LogIntoGoogle runs the Log into Google action.
func (c *AuthenticationController) LogIntoGoogle(ctx *app.LogIntoGoogleAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
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
	if err != nil {
		fmt.Println("Issue parsing response body:", err)
		return nil
	}
	fmt.Fprintf(w, "{Content: %s, AuthToken: %s}", contents, token.AccessToken)
	return nil
}

// LogIntoFacebook runs the Log into Google action.
func (c *AuthenticationController) LogIntoFacebook(ctx *app.LogIntoFacebookAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request
	fbRedirectURL := &url.URL{
		Scheme: "http",
		Path:   facebook.Endpoint.AuthURL,
	}
	queries := fbRedirectURL.Query()
	queries.Set("app_id", facebookOauthConfig.ClientID)
	queries.Set("redirect_uri", facebookOauthConfig.RedirectURL)
	queries.Set("state", oauthStateString)
	fbRedirectURL.RawQuery = queries.Encode()
	goa.Info(ctx, "redirect", goa.KV{Key: "EscapedPath", Value: fbRedirectURL.RequestURI()})
	http.Redirect(w, r, fbRedirectURL.RequestURI(), http.StatusTemporaryRedirect)
	return nil
}

// CallbackResponseFromFacebook runs the Callback response from Facebook action.
func (c *AuthenticationController) CallbackResponseFromFacebook(ctx *app.CallbackResponseFromFacebookAuthenticationContext) error {
	w := goa.Response(ctx).ResponseWriter
	r := goa.Request(ctx).Request

	code := r.FormValue("code")
	state := r.FormValue("state")
	// fmt.Fprintf(w, "Hello! Code is: %s", code)

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}
	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	fbTokenURL := &url.URL{
		Scheme: "http",
		Path:   facebook.Endpoint.AuthURL,
	}
	queries := fbTokenURL.Query()
	queries.Set("client_id", facebookOauthConfig.ClientID)
	queries.Set("redirect_uri", facebookOauthConfig.RedirectURL)
	queries.Set("client_secret", secrets.Facebook.Secret)
	queries.Set("code", code)
	fbTokenURL.RawQuery = queries.Encode()

	response, err := http.Get(fbTokenURL.RequestURI())
	if err != nil {
		return err
	}
	defer response.Body.Close()
	goa.Info(ctx, "token received", goa.KV{Key: "AccessToken", Value: token.AccessToken})

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
