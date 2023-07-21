package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the JSON file
	filePath := "./logs/new.json"
	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read the JSON file: %v\n", err)
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v\n", err)
		return
	}

	// actually format the json
	formattedJSON, _ := json.MarshalIndent(jsonData, "", "  ")

	// this puts the json into the file new.json
	err = ioutil.WriteFile("./logs/new.json", formattedJSON, 0644)
}