package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, upperCase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
}
