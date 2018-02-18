package main

import "fmt"

func main() {
	//reference
	reference := []int{1, 2, 3, 4, 5, 6}

	example1 := reference[2:4]
	fmt.Println("Printing example1 slice:", example1)
	fmt.Println("length of the slice:", len(example1))
	fmt.Println("capacity of the slice:", cap(example1))
	fmt.Println("what slice can really see:", example1[:cap(example1)])

	fmt.Println("###############")

	example2 := make([]int, len(reference[2:4]), cap(reference[2:4]))
	copy(example2, reference[2:4])
	fmt.Println("Printing example2 slice:", example2)
	fmt.Println("length of the slice:", len(example2))
	fmt.Println("capacity of the slice:", cap(example2))
	fmt.Println("what slice can really see:", example2[:cap(example2)])

	fmt.Println("###############")

	fmt.Println("changing example1[0] = 55")
	example1[0] = 55
	fmt.Println("Printing reference slice:", reference)
	fmt.Println("Printing example1 slice:", example1)
	fmt.Println("Printing example2 slice:", example2) //this is not changed, because it cointains different underlying array

	fmt.Println("###############")
	capTest := reference[1:3:3] // sets capacity to 2
	fmt.Println("Printing capTest slice:", capTest)
	fmt.Println("Printing capacity capTest slice:", cap(capTest))

}
