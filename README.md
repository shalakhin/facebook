# facebook

Facebook Graph API for Go.

It is work in progress for now and pull requests are highly appreciated.

## Installation & Update

```bash
$ go get github.com/OShalakhin/facebook # install
$ go get -u github.com/OShalakhin/facebook # update
```

## Usage

```go
// Example
var graph := facebook.New(
        "AppID",
        "Secret",
        "https://example.com/facebook/callback",
        []string{"email"},
)

// Signup redirects user to facebook
func Signup(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, graph.AuthURL(""), http.StatusFound)
}

// Handle response
func HandleOAuth(w http.ResponseWriter, r *http.Request) {
        graph.GetAccessToken(r)
        // now you can get access token
        fmt.Fprintf(w, graph.AccessToken)
}
```

## Docs

[Godoc](http://godoc.org/github.com/OShalakhin/facebook)

## TODO

- error handling
- tests coverage
- graph API:
  - debug token
  - get user info
  - get user image
  - message user
  - ...
