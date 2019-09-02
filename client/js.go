package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var data = []byte(`
    {
        "data": {
                "username": "admin",
                "password": "admin"
        }
    }
`)

type Inn struct {
	Username   string   `json:"username"`
	Password  string   `json:"password"`
}

type Info struct {
	Data Inn `json:"data"`
}

func main() {

	var info Info
	if err := json.Unmarshal(data, &info); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", info)

}
