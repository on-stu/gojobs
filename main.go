package main

import "fmt"

func main() {
	minsu := map[string]string{"name": "minsu", "age": "12"}
	fmt.Println(minsu)
	for key, value := range minsu {
		fmt.Println(key, value)
	}
}

// map은 약간 object나 dict와 비슷한건데 정확히는 다름.
// 왜냐하면 모든 value의 타입이 같아야함
