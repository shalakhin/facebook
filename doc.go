// Package facebook implements Facebook Graph API.
//
// request token url is: https://www.facebook.com/dialog/oauth
// access token url is: https://graph.facebook.com/oauth/access_token
//
// These urls are hidden inside Graph struct as they don't change.
//
//     // Example
//     var graph := facebook.New(
//             "AppID",
//             "Secret",
//             "https://example.com/facebook/callback",
//             []string{"email"},
//     )
//
//     // Signup redirects user to facebook
//     func Signup(w http.ResponseWriter, r *http.Request) {
//             http.Redirect(w, r, graph.AuthURL(""), http.StatusFound)
//     }
//
//     // Handle response
//     func HandleOAuth(w http.ResponseWriter, r *http.Request) {
//             var g *facebook.Graph
//             g = graph
//             g.GetAccessToken(r)
//     }
//
package facebook
