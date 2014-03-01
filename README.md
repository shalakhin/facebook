# facebook

Facebook Graph API fo Go. It is work in progress for now and pull requests
are highly appreciated.


```go
// initialize graph
Graph := facebook.New(key, secret)
// get request token
Graph.RequestToken(state, scope)
// get access token
Graph.GetAccess(r)
// is it authenticated
Graph.IsAuthenticated()
// API request to the
Graph.API(URL, map[string]string)
Graph.User.Get()
Graph.Image.Get()
```
