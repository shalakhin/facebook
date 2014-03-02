// Package facebook implements Facebook Graph API.
//
// request token url is: https://www.facebook.com/dialog/oauth
// access token url is: https://graph.facebook.com/oauth/access_token
//
// These urls are hidden inside Graph struct as they don't change.
//
//     // Example
//     graph, err := facebook.New("mykey", "mysecret")
//     // ... handle error case
//     // get request token
//     graph.Authenticate(scope)
package facebook
