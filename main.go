package main

import "fmt"

func canIDrink(age int) bool {
	switch { //여기도 variable expression을 사용 가능
	case age < 18: //else-if 를 남발하지 않기 위해 이렇게도 사용 가능
		return true
	case age == 18:
		return true
	default:
		return false
	}

}

func main() {
	result := canIDrink(17)
	fmt.Println(result)
}
