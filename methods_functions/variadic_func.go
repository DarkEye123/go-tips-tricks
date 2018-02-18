package main

import "fmt"

func variadicExample(nums ...int) int {
	var tmp int
	for _, val := range nums {
		tmp += val
	}
	return tmp
}

func main() {
	fmt.Println(variadicExample(1, 2, 3, 4, 5))
	variadingParam := []int{1, 2, 3, 4, 5}
	fmt.Println(variadicExample(variadingParam...))
}
