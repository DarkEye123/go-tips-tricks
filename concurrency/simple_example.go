package main

import "fmt"

func main() {
	c := make(chan string)
	go hiFiveNtimes(c, 3)
	for s := range c {
		fmt.Println(s)
	}
	if _, ok := <-c; !ok {
		fmt.Println("Channel was closed")
	}
}

func hiFiveNtimes(c chan string, times int) {
	for x := 0; x < times; x++ {
		c <- "hi"
	}
	close(c)
}
