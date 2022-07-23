package main

import (
	"fmt"

	"github.com/on-stu/gojobs/dictionary"
)

func main() {
	myDictionary := dictionary.Dictionary{"secondd": "2"}
	myDictionary.Add("second", "wow")
	definition, err := myDictionary.Search("second")
	if err == nil {
		fmt.Println(definition)
	} else {
		fmt.Println("error")
	}
}
