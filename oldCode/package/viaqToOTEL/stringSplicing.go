package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	// changing time
	oldTime := "2022-10-20T15:11:30.764362932Z"
	parsedTime, _ := time.Parse(time.RFC3339Nano, oldTime)
	unixNanoTime := parsedTime.UnixNano()
	fmt.Println(unixNanoTime)

	// string splicing
	oldImage := "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ed90fd557cc619f98a99bc8c552ee7b8a8ee369a3a2cdf2f9a4727878d2d049e"
	// this splices the string from the beginning to the first : and then puts everything after that into a new string
	newImage := oldImage[strings.Index(oldImage, ":"):]
	fmt.Println(newImage)
}
