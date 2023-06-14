package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s1 := Student{
		Name: "Sagar",
		Age:  23,
	}

	yamlData, err := yaml.Marshal(&s1)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fmt.Println(" --- YAML ---")
	fmt.Println(string(yamlData)) // yamlData will be in bytes. So converting it to string.
}
