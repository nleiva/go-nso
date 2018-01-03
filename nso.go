package main

import (
	"net/http"
	"net/url"
	"strings"
)

func fullConfig(u *url.URL, d string) (req *http.Request, err error) {
	// "http://admin:admin@mrstn-nso.cisco.com:8080/api/config/devices/device/mrstn-5501-1.cisco.com/config?deep=true"
	u.Path = "api/config/devices/device/" + d + "/config"
	// All the details
	q := u.Query()
	q.Set("deep", "true")
	u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/vnd.yang.data+json")
	return req, err
}

func interfaceConfig(u *url.URL, d string) (req *http.Request, err error) {
	u.Path = "api/running/devices/device/" + d + "/config/interface/TenGigE"
	// All the details
	q := u.Query()
	q.Set("deep", "true")
	u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/vnd.yang.collection+json")
	return req, err
}

func routerConfig(u *url.URL, d string) (req *http.Request, err error) {
	// static, isis, bgp, etc...
	p := "static"
	u.Path = "api/running/devices/device/" + d + "/config/router/" + p
	// All the details
	q := u.Query()
	q.Set("deep", "true")
	u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/vnd.yang.data+json")
	return req, err
}

func setRouterConfig(u *url.URL, d string, p string, c string) (req *http.Request, err error) {
	// PUT, POST -> REPLACE. PATH -> MERGE
	// p = static, isis, bgp, etc...
	u.Path = "api/running/devices/device/" + d + "/config/router/" + p
	// Request
	req, err = http.NewRequest("PATCH", u.String(), strings.NewReader(c))
	req.Header.Add("Content-Type", "application/vnd.yang.data+json")
	req.Header.Add("Accept", "application/vnd.yang.data+json")
	return req, err
}

func syncFrom(u *url.URL, d string) (req *http.Request, err error) {
	u.Path = "api/running/devices/device/" + d + "/_operations/sync-from"
	// Request
	req, err = http.NewRequest("POST", u.String(), nil)
	req.Header.Add("Content-Type", "application/vnd.yang.operation+json")
	req.Header.Add("Accept", "application/vnd.yang.operation+json")
	return req, err
}
