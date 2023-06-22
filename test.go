package main

import(
	"fmt"
	"time"
	"strconv"
)

func main() {
	oldTime := "2022-10-20T15:11:30.764362932Z"
	parsedTime, _ := time.Parse(time.RFC3339Nano, oldTime)
	unixNanoTime := parsedTime.UnixNano()
	fmt.Print(strconv.Itoa(int(unixNanoTime)))
}