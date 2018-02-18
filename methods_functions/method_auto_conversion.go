package main

import "fmt"

type TestS int

func (t *TestS) Add(val int) {
	*t = *t + TestS(val)
}

func main() {
	var t TestS = 10
	t.Add(10) // go converts it to something like this (&t).Add(10)
	fmt.Println(t)
}
