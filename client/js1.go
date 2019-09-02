package main
 
import (
    "encoding/json"
    "fmt"   
)
 
func main() {
    jStr := `
    {
        "AAA":{
            "BBB": {
                "CCC": ["1111"],
                "DDD": ["2222", "3333"]
            }
        }
    }
    `
 
    type Inner struct {
        Key2 []string `json:"CCC"`
        Key3 []string `json:"DDD"`
    }
    type Outer struct {
        Key Inner `json:"BBB"`
    }
    type Outmost struct {
        Key Outer `json:"AAA"`
    }
    var cont Outmost
    json.Unmarshal([]byte(jStr), &cont)        
    fmt.Printf("%+v\n", cont)
}
