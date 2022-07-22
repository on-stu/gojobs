package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	minsu := person{name: "minsu", age: 18}
	fmt.Println(minsu.age)
}
