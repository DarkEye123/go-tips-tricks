package main

import "fmt"

type Empty struct{}

func main() {
	reference := []*Empty{&Empty{}, &Empty{}, &Empty{}, &Empty{}} //imagine this could be like 10 thousand items stored here
	buggy := buggySlice(reference)                                // returns first two pointers
	fmt.Println("comparation of what reference and buggy can see")
	fmt.Println(reference[:cap(reference)], "vs", buggy[:cap(buggy)])

	fmt.Println("##############")

	correct := correctSlice(reference) // returns first two pointers
	fmt.Println("comparation of what reference and correct can see")
	fmt.Println(reference[:cap(reference)], "vs", correct[:cap(correct)])
}

// returns same underlying array, therefore garbage collector can't free the space
func buggySlice(reference []*Empty) []*Empty {
	return reference[:2] // first two pointers
}

// new underlying array allocated via copy, therefore reference can be freed correctly in the code after usage
func correctSlice(reference []*Empty) []*Empty {
	correct := make([]*Empty, 2)
	copy(correct, reference[:2]) // first two pointers
	return correct
}
