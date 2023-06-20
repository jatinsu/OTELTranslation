package main

import (
	"fmt"
	"time"
)

func main() {
	oldTime := "2022-10-20T15:11:30.764362932Z"
	parsedTime, _ := time.Parse(time.RFC3339Nano, oldTime)
	unixNanoTime := parsedTime.UnixNano()
	fmt.Println(unixNanoTime)
}
