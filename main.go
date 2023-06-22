package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"workspace/logconverter"

	"github.com/openshift/cluster-logging-operator/test/helpers/types"
)

func main() {
	// this imports the json file and puts in logJson
	logJson, err := ioutil.ReadFile("Logs/viaq.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var log types.ContainerLog
	json.Unmarshal([]byte(logJson), &log)

	//fmt.Println(string(logJson))
	//Update the function call to reference the logconverter package
	theNewLog := logconverter.ConvertLog(log)

	outputJSON, _ := json.MarshalIndent(theNewLog, "", "    ")
	fmt.Println(string(outputJSON))
	ioutil.WriteFile("Logs/newLog.json", []byte(outputJSON), 0644)
}
