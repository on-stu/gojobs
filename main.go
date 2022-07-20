package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpperNaked(name string) (length int, uppercase string) { //naked function이라고 하는건데 return을 이 줄에서 바로 해줌
	length = 1
	uppercase = "HI"
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
	repeatMe("hi", "hello", "wow")
}
