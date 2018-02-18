package main

import "fmt"

type TestS int

func (t *TestS) Add(val int) {
	*t = *t + TestS(val)
}

func main() {
	t := new(TestS)
	t.Add(10) // 't' is copied by reference, therefore this instance will be added a value '10'
	fmt.Println(*t)
}
