# go-nso

Playing around with NSO REST-API and Go. This is work in progress, will make it more user-friendly soon.

## Configure static routes

```go
u := new(url.URL)
u.Scheme = "http"
u.Host = "mrstn-nso.cisco.com:8080"
u.User = url.UserPassword("admin", "admin")
device := "mrstn-5501-1.cisco.com"

var netClient = &http.Client{
	Timeout: time.Second * 20,
}

// Generate the message payload
config, err := generateStatic("2001:420::/32", "2001:420:2cff:1204::1")
...

// Use NSO REST-API to configure the route
req, err := setRouterConfig(u, device, "static", config)
...
resp, err := netClient.Do(req)
```

## Sync from device

```go
u := new(url.URL)
u.Scheme = "http"
u.Host = "mrstn-nso.cisco.com:8080"
u.User = url.UserPassword("admin", "admin")
device := "mrstn-5501-1.cisco.com"

var netClient = &http.Client{
	Timeout: time.Second * 20,
}

// Use NSO REST-API to Sync from device
req, err := syncFrom(u, device)
...
resp, err := netClient.Do(req)
```

## Request device config

```go
u := new(url.URL)
u.Scheme = "http"
u.Host = "mrstn-nso.cisco.com:8080"
u.User = url.UserPassword("admin", "admin")
device := "mrstn-5501-1.cisco.com"

var netClient = &http.Client{
	Timeout: time.Second * 20,
}

// Use NSO REST-API to Sync from device
req, err := routerConfig(u, device)
...
resp, err := netClient.Do(req)
```