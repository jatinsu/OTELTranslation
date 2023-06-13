package main

import (
	"encoding/json"
	"fmt"
)


type Bird struct {
	Species string
	Description string
	Dimensions Dimensions
}

type Dimensions struct {
	Height int
}



func main(){

	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var birds Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Println(birds.Description)
	// {pigeon likes to perch on rocks {24 10}}

}