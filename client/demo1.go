// Copyright (c) 2012-2013 Jason McVetta.  This is Free Software, released
// under the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for
// details.  Resist intellectual serfdom - the ownership of ideas is akin to
// slavery.

// Example demonstrating use of package restclient, with HTTP Basic
// authentictation over HTTPS, to retrieve a Github auth token.
package main

/*

NOTE: This example may only work on *nix systems due to gopass requirements.

*/

import (
	"encoding/json"
	"fmt"
	"restclient"
	"log"

)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

var data1 = []byte(`
    {
        "data": {
                "username": "admin",
                "password": "admin"
        }
    }
`)

var data = []byte(` { "data": { "username": "admin", "password": "admin" } } `)

type Inn struct {
        Username   string   `json:"username"`
        Password  string   `json:"password"`
}

type Info struct {
        Data Inn `json:"data"`
}

type Resp struct {
	AppName string `json:"app_name"`
	CreationTime int `json:"creation_time"`
	ExpiryTime int `json:"expiry_time"`
	Id string `json:"id"`
	LastModified int `json:"last_modified"`
	SessionToken string `json:"session_token"`
	SourceIp string `json:"source_ip"`
	Username string  `json:"username"`
}

type Rinfo struct {
        Data Resp `json:"data"`
}

func main() {
        var info Info
        if err := json.Unmarshal(data, &info); err != nil {
                log.Fatal(err)
        }
	fmt.Printf("Sending data: %+v\n", info)

	//
	// Struct to hold response data
	//
	res:= Rinfo{}

	//
	// Struct to hold error response
	//
	e := struct {
		Message string
	}{}

	//
	// Setup HTTP Basic auth (ONLY use this with SSL)
	//
	rr := restclient.RequestResponse{
		Url:      "https://10.18.151.28:5392/v1/tokens",
		Method:   "POST",
		Data:     &info,
		Result:   &res,
		Error:    &e,
	}

	//
	// Send request to server
	//
	status, err := restclient.Do(&rr)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Process response
	//
	println("GMD proxy got response: \n")
	if status == 201 {
		fmt.Printf("Got auth data: %s\n", res.Data)
		fmt.Printf("Got auth ID: %s\n", res.Data.Id)
		fmt.Printf("Got auth token: %s\n", res.Data.SessionToken)
		fmt.Printf("Got auth app_name: %s\n", res.Data.AppName)
		fmt.Printf("Got auth creation_time: %d\n", res.Data.CreationTime)
		fmt.Printf("Got auth expiry_time: %d\n", res.Data.ExpiryTime)
		fmt.Printf("Got auth LastModified: %d\n", res.Data.LastModified)
		fmt.Printf("Got auth source_ip: %s\n", res.Data.SourceIp)
		fmt.Printf("Got auth Username: %s\n\n", res.Data.Username)

		fmt.Printf("**** Will use auth token: %s ****\n", res.Data.SessionToken)
	} else {
		fmt.Println("Bad response status from Array")
		fmt.Printf("\t Status:  %v\n", status)
		fmt.Printf("\t Status:  %v\n", res)
		fmt.Printf("\t Message: %v\n", e.Message)
	}
	println("")
}
