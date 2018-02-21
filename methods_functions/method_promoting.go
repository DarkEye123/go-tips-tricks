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

func main() {
	x := NewEmbededRand(55)
	fmt.Println(x.Intn(3)) //method promoting
}
