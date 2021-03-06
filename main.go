package main

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
)

var jsonData = `

{ 
    "name": "John", 
    "age": 31, 
    "city": "New York"
}

`

//Person is struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

var data = []byte(jsonData)

func main() {

	var person Person
	json.Unmarshal(data, &person)

	fmt.Println(person)

	value, typevalue, offset, _ := jsonparser.Get(data, "name")

	fmt.Println(string(value), typevalue, offset)

}
