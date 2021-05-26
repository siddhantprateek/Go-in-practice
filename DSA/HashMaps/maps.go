package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {

	var person = map[string]string{
		"name": "john doe",
		"age":  "23",
	}

	var profile = make(map[string]string)
	profile["name"] = "jane doe"
	profile["age"] = "23"

	var user = make(map[string]interface{})

	user["name"] = "peter parker"
	user["age"] = 30

	var users = []map[string]interface{}{
		{"name": "monkey d lufy", "age": 19},
		{"name": "trafagar d law", "age": 23},
		{"name": "nico robin", "age": 20},
	}

	var userStruct = map[string]User{
		"name": {Name: "Monkey D Lufy"},
		"age":  {Age: 19},
	}

	fmt.Println(person)
	fmt.Println(profile)
	fmt.Println(user)
	fmt.Println(users)
	fmt.Println(userStruct)

}