package main

import "fmt"

func main() {
	names := []string{"nico", "lynn", "dal"} //여기 길이를 지정해주면 array, 아니면 slice가 된다.
	names = append(names, "wow")
	fmt.Println(names)
}

//&을 입력하면 메모리 주소를 줌
