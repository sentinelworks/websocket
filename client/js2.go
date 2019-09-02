package main
 
import (
    "encoding/json"
    "fmt"   
)
 
func main() {
    jStr := `
    {
            "BBB": {
                "CCC": "1111"
                "DDD": "2222" 
            }
        }
    `
 
    type Inner struct {
        Key2 string `json:"CCC"`
        Key3 string `json:"DDD"`
    }
    type Outmost struct {
        Key Inner `json:"BBB"`
    }
    var cont Outmost
    json.Unmarshal([]byte(jStr), &cont)        
    fmt.Printf("%+v\n", cont)
}
