package main

import "fmt"

func main() {
	minsu := map[string]string{"name": "minsu", "age": "12"}
	fmt.Println(minsu)
	for key, value := range minsu {
		fmt.Println(key, value)
	}
}
