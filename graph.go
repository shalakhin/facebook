package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type (
	// Graph is a core of this package as its methods allow to communicate with
	// Facebook Graph API
	Graph struct {
		// Register your app at https://developers.facebook.com to get AppID and Secret
		AppID       string
		Secret      string
		AccessToken string
		Expire      time.Duration
		// Full list of scope options here:
		// https://developers.facebook.com/docs/facebook-login/permissions/
		Scope []string

		apptoken        string
		requestTokenURL *url.URL
		accessTokenURL  *url.URL
		callbackURL     *url.URL
	}
	// DebugInfo holds information from debug info on particular token
	DebugInfo struct {
		Data DebugData `json:"data"`
	}
	// DebugData contains all fields returned in DebugInfo
	DebugData struct {
		AppID       int       `json:"app_id"`
		Application string    `json:"application"`
		ExpiresAt   EpochTime `json:"expires_at"`
		IsValid     bool      `json:"is_valid"`
		IssuedAt    EpochTime `json:"issued_at,omitempty"`
		Scopes      []string  `json:"scopes"`
		UserID      int       `json:"user_id"`
	}
)

// New initializes Graph instance
func New(appID, secret, callback string, scope []string) *Graph {
	var reqTok, accessTok, callbackURL *url.URL
	reqTok, _ = url.Parse("https://www.facebook.com/dialog/oauth")
	accessTok, _ = url.Parse("https://graph.facebook.com/oauth/access_token")
	callbackURL, _ = url.Parse(callback)
	return &Graph{
		AppID:           appID,
		Secret:          secret,
		Scope:           scope,
		apptoken:        strings.Join([]string{appID, secret}, "|"),
		requestTokenURL: reqTok,
		accessTokenURL:  accessTok,
		callbackURL:     callbackURL,
	}
}

// AuthURL generates URL to redirect to. User will give permission and you'll recieve
// request token. You can pass state parameter to protect from the CSRF
func (g *Graph) AuthURL(state string) string {

	query := g.requestTokenURL.Query()
	query.Set("client_id", g.AppID)
	query.Set("redirect_uri", g.callbackURL.String())
	query.Set("scope", strings.Join(g.Scope, ","))
	query.Set("response_type", "code")
	if state != "" {
		query.Set("state", state)
	}
	g.requestTokenURL.RawQuery = query.Encode()
	return g.requestTokenURL.String()
}

// Authenticate parses request for code and retrieve access token from
// response and expiration. In case of errors returns error which you can
// handle on your own (i.e. redirect with error message or return 500 page or
// else.
func (g *Graph) Authenticate(r *http.Request) error {
	var err error
	var resp *http.Response
	var result []byte
	var expire time.Duration
	var values url.Values

	query := g.accessTokenURL.Query()
	query.Set("client_id", g.AppID)
	query.Set("redirect_uri", g.callbackURL.String())
	query.Set("client_secret", g.Secret)
	query.Set("code", r.URL.Query().Get("code"))
	g.accessTokenURL.RawQuery = query.Encode()

	if resp, err = http.Get(g.accessTokenURL.String()); err != nil {
		return err
	}
	defer resp.Body.Close()

	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	if values, err = url.ParseQuery(string(result)); err != nil {
		return err
	}
	durationStr := fmt.Sprintf("%ss", values.Get("expires"))
	if expire, err = time.ParseDuration(durationStr); err != nil {
		return err
	}

	g.AccessToken = values.Get("access_token")
	g.Expire = expire
	return nil
}

// DebugToken retrieves debug information to verify token. Look here:
// https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow#checktoken
func (g *Graph) DebugToken(token string) (*DebugInfo, error) {
	var resp *http.Response
	var err error

	info := DebugInfo{}
	if token == "" {
		return &info, errors.New("empty token")
	}

	debugURL, _ := url.Parse("https://graph.facebook.com/debug_token")
	query := debugURL.Query()
	query.Set("input_token", token)
	query.Set("access_token", g.apptoken)
	debugURL.RawQuery = query.Encode()

	if resp, err = http.Get(debugURL.String()); err != nil {
		return &info, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return &info, err
	}
	return &info, nil
}

// Debug access token received after Authenticate
func (g *Graph) Debug() (*DebugInfo, error) {
	return g.DebugToken(g.AccessToken)
}
