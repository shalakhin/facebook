# facebook [![GoDoc](https://godoc.org/github.com/OShalakhin/facebook?status.png)](https://godoc.org/github.com/OShalakhin/facebook)

Facebook Graph API for Go.

[Documentation on Godoc](http://godoc.org/github.com/OShalakhin/facebook)

Implemented:

- OAuth supported
- Debug information
- User information (partially)
- User picture

It is work in progress for now and pull requests are highly appreciated.

## Installation & Update

```bash
$ go get github.com/OShalakhin/facebook # install
$ go get -u github.com/OShalakhin/facebook # update
```

## Usage

```go
// Example
package main

import (
    "fmt"
    "net/http"

    "github.com/OShalakhin/facebook"
)

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
        var g *facebook.Graph
        g = graph
        g.GetAccessToken(r)
        // now you can get access token
        fmt.Fprintf(w, graph.AccessToken)
}
```
