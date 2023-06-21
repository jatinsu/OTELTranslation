package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Key1 string `yaml:"key1"`
	Key2 string `yaml:"key2"`
}

func main() {
	// Read the YAML file
	yamlFile, err := ioutil.ReadFile("parsing/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// Unmarshal the YAML into a struct
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Update the struct with desired values
	config.Key1 = "Value1"
	config.Key2 = "Value2"

	// Marshal the struct back to YAML
	updatedYAML, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}

	// Write the updated YAML to a file or use it as needed
	err = ioutil.WriteFile("updated_config.yaml", updatedYAML, 0644)
	if err != nil {
		log.Fatalf("Failed to write updated YAML file: %v", err)
	}

	fmt.Println("Updated YAML file has been created.")
}
