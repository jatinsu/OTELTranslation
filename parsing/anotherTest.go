package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"

)

type Data struct {
	Spec     Spec     `json:"spec"`
	Metadata Metadata `json:"metadata"`
}

type Spec struct {
	List      []string          `json:"list"`
	NestedMap map[string]string `json:"nestedMap"`
}

type Metadata struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func main() {
	inputJSON := `
	{
	  "metadata": {
	    "name": "example",
	    "version": "1.0"
	  },
	  "spec": {
	    "list": ["value1", "value2", "value3"],
	    "nestedMap": {
	      "key3": "value3",
	      "key1": "value1",
	      "key2": "value2"
	    }
	  }
	}`

	var data Data
	err := json.Unmarshal([]byte(inputJSON), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	outputData := Data{
		Spec:     data.Spec,
		Metadata: data.Metadata,
	}

	outputJSON, err := json.Marshal(outputData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	formattedJSON := pretty.Pretty(outputJSON)

	fmt.Println(string(formattedJSON))
}
