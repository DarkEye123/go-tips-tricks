package main

import "fmt"

type TestS struct {
	val int
}

func (t TestS) Add(val int) {
	t.val += val
}

func main() {
	t := TestS{}
	t.Add(10) // 't' is copied by value, therefore this instance won't be added a value '10'
	fmt.Println(t.val)
}
