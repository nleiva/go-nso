package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	u := new(url.URL)
	u.Scheme = "http"
	u.Host = "mrstn-nso.cisco.com:8080"
	u.User = url.UserPassword("admin", "admin")

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	device := "mrstn-5501-1.cisco.com"

	//req, err := fullConfig(u, device")
	//req, err := interfaceConfig(u, device)
	//req, err := routerConfig(u, device)
	//req, err := syncFrom(u, device)
	//config, err := generateStatic("191.0.0.0/8", "10.87.89.1")
	config, err := generateStatic("2001:425::/32", "2001:420:2cff:1204::1")
	checkErr(err)

	req, err := setRouterConfig(u, device, "static", config)
	checkErr(err)

	resp, err := netClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	// Read JSON data
	/* 	data := new(Router)
	   	err = decodeJSON(data, resp.Body)
	   	checkErr(err)
	   	fmt.Printf("%v\n", data)
	*/

	contents, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Printf("%s\n", string(contents))
}
