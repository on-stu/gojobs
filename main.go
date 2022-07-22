package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 200
	fmt.Println(a, *b)
}

//&을 입력하면 메모리 주소를 줌
