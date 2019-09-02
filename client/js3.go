package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var data = []byte(`
    {
        "client": {
                "id": "2212fw",
                "name": "Papupapa Hernandez",
                "email": "papupapa@gmail.com",
                "phones": ["554-223-2311", "332-232-2123"]
        }
    }
`)

type Client struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Phones []string `json:"phones"`
}

type Info struct {
	Key Client `json:"client"`
}

func main() {

	var info Info
	if err := json.Unmarshal(data, &info); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", info)

}
