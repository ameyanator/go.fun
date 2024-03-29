package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Parsing JSON without knowing the structure
//How?
//Parse JSON into an interface{} instead of a struct.
//Then inspect it and use

var ks = []byte(`{
	"firstName": "Jean",
	"lastName": "Bartik",
	"age": 86,
	"education": [
	{
	"institution": "Northwest Missouri State Teachers College",
	"degree": "Bachelor of Science in Mathematics"
	},
	{
	"institution": "University of Pennsylvania",
	"degree": "Masters in English"
	}
	],
	"spouse": "William Bartik",
	"children": [
	"Timothy John Bartik",
	"Jane Helen Bartik",
	"Mary Ruth Bartik"
	]
	}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f)
	m := f.(map[string]interface{})
	fmt.Println(m["education"])
	fmt.Println("Printing JSON")
	printJSON(f)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is array")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is object")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown Type")
	}
}
