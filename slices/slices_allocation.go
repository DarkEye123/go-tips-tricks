package main

import "fmt"

func main() {
	//reference
	reference := []int{1, 2, 3, 4, 5, 6}

	example1 := reference[2:4]
	// content of the slice
	fmt.Println("Printing example1 slice:", example1)
	// length of the slice
	fmt.Println("length of the slice:", len(example1))
	// what is the capacity of the slice?
	fmt.Println("capacity of the slice:", cap(example1))
	// what is trully allocated? What an slice can really see
	fmt.Println("what slice can really see:", example1[:cap(example1)])

	fmt.Println("###############")

	example2 := make([]int, len(reference[2:4]), cap(reference[2:4]))
	copy(example2, reference[2:4])
	// content of the slice
	fmt.Println("Printing example2 slice:", example2)
	// length of the slice
	fmt.Println("length of the slice:", len(example2))
	// what is the capacity of the slice?
	fmt.Println("capacity of the slice:", cap(example2))
	// what is trully allocated? What an slice can really see
	fmt.Println("what slice can really see:", example2[:cap(example2)])
	fmt.Println("###############")

	fmt.Println("changing example1[0] = 55")
	example1[0] = 55

	fmt.Println("Printing reference slice:", reference)
	fmt.Println("Printing example1 slice:", example1)
	fmt.Println("Printing example2 slice:", example2) //this is not changed, because it cointains different underlying array
}
