package main

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Println("Please enter the name")
	fmt.Scanln(&name)

	fmt.Println("Please enter the address")
	fmt.Scanln(&address)

	m := map[string]string{
		"name":    name,
		"address": address,
	}

	var json, err = json.Marshal(m)

	if err == nil {
		fmt.Println(string(json))
	} else {
		fmt.Println("Error: ", err)
	}
}
