package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
	"workspace/logconverter"
)

func TestConvertLog(t *testing.T) {
	// Read the sample JSON log file
	logJson, err := ioutil.ReadFile("Logs/newLog.json")
	if err != nil {
		t.Errorf("Failed to read sample log file: %v", err)
		return
	}

	var log logconverter.Log
	err = json.Unmarshal(logJson, &log)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
		return
	}

	// Convert the log
	newLog := logconverter.ConvertLog(log)

	// Convert the new log to JSON
	outputJSON, err := json.Marshal(newLog)
	if err != nil {
		t.Errorf("Failed to marshal new log to JSON: %v", err)
		return
	}

	// Read the expected JSON file
	expectedJSON, err := ioutil.ReadFile("Logs/otel.json")
	if err != nil {
		t.Errorf("Failed to read expected log file: %v", err)
		return
	}

	// Compare the overall JSON structures using DeepEqual
	if !reflect.DeepEqual(outputJSON, expectedJSON) {
		t.Error("Converted log JSON does not match the expected JSON structure")
	}

	// Optionally, you can also write the output JSON to a file for inspection if needed
	err = ioutil.WriteFile("Logs/testLog.json", outputJSON, 0644)
	if err != nil {
		t.Errorf("Failed to write converted log to file: %v", err)
	}
}
