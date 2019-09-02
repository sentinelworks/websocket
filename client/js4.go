package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var data = []byte(`
    {
        "Data": {
                "id": "2212fw",
                "email": "papupapa@gmail.com"
        }
    }
`)

type Client struct {
	Id     string   `json:"id"`
	Email  string   `json:"email"`
}

type Info struct {
	data Client `json:"Data"`
}

func main() {

	var info Info
	if err := json.Unmarshal(data, &info); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", info)

}
