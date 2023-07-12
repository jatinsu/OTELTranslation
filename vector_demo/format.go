package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func format() {
	// Read the JSON file
	filePath := "new.json"
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
	err = ioutil.WriteFile("new.json", formattedJSON, 0644)
}

func main() {
	counter:=0

	for counter < 100000 {
		time.Sleep(2 * time.Second)	
		counter++
		format()
	}
}