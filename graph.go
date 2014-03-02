package facebook

import (
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

		requestTokenURL *url.URL
		accessTokenURL  *url.URL
		callbackURL     *url.URL
	}
)

// New initializes Graph instance
func New(appID, secret, callback string, scope []string) *Graph {
	var reqTok, accessTok, callbackURL *url.URL
	reqTok, _ = url.Parse("https://www.facebook.com/dialog/oauth")
	accessTok, _ = url.Parse("https://graph.facebook.com/oauth/access_token")
	callbackURL, _ = url.Parse(callback)
	return &Graph{
		AppID:  appID,
		Secret: secret,
		Scope:  scope,
		// AccessToken: "",
		// Expire: time.Time{},
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

// GetAccessToken parses request for code and retrieve access token from
// response and expiration. In case of errors returns error which you can
// handle on your own (i.e. redirect with error message or return 500 page or
// else.
func (g *Graph) GetAccessToken(r *http.Request) error {
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
	if expire, err = time.ParseDuration(values.Get("expires")); err != nil {
		return err
	}

	g.AccessToken = values.Get("access_token")
	g.Expire = expire
	return nil
}
