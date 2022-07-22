package main

import (
	"fmt"

	"github.com/on-stu/gojobs/dictionary"
)

func main() {
	dict := dictionary.Dictionary()
	fmt.Println(dict)
}
