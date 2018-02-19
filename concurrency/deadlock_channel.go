package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	ch <- "again"
	fmt.Println(<-ch, <-ch)

}
