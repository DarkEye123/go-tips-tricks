package main

import "fmt"

type Empty struct {
	name string
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6}

	//remove one item, '2'
	x = append(x[:1], x[2:]...) // cut the "2"
	fmt.Println(x)
	fmt.Println(x[:cap(x)])

	fmt.Println("#############")

	// remove '4' and '5'
	// currently x = [1 3 4 5 6]
	x = append(x[:2], x[4:]...) // cut the "2"
	fmt.Println(x)
	fmt.Println(x[:cap(x)])

	fmt.Println("#############")

	// ###################################

	// pointers demonstration (cut, delete)
	r := []*Empty{&Empty{name: "a"}, &Empty{name: "b"}, &Empty{name: "c"}, &Empty{name: "d"}} //imagine this could be like 10 thousand items stored here

	//cut 'b' and 'c'
	copy(r[1:], r[3:]) // r = [a d c d]
	for k, n := len(r)-3+1, len(r); k < n; k++ {
		r[k] = nil
	} // r = [a d nil nil]
	r = r[:len(r)-3+1]
	for _, f := range r {
		fmt.Println(f.name)
	}

	fmt.Println("#############")
	r = []*Empty{&Empty{name: "a"}, &Empty{name: "b"}, &Empty{name: "c"}, &Empty{name: "d"}} //imagine this could be like 10 thousand items stored here

	//delete only one item - 'b'
	r, r[len(r)-1] = append(r[:1], r[2:]...), nil
	for _, f := range r {
		fmt.Println(f.name)
	}

}
