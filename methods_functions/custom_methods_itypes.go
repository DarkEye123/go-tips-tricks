package main

import (
	"fmt"
	"math/rand"
)

type EmbededRand struct {
	*rand.Rand
	val int
}

func NewEmbededRand(seed int64) (ret *EmbededRand) {
	ret = new(EmbededRand)
	ret.Rand = rand.New(rand.NewSource(seed))
	return
}

func (e *EmbededRand) Intn(x int) int {
	fmt.Print("Outer method ")
	return e.Rand.Intn(x)
}

func main() {
	x := NewEmbededRand(55)
	fmt.Println(x.Intn(3))
	fmt.Println("Rand meethod", x.Rand.Intn(3))
}
