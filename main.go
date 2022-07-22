package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 { //조건문 내부에서 변수 생성 가능
		//variable expression
		return false
	}
	return true

}

func main() {
	result := canIDrink(17)
	fmt.Println(result)
}
