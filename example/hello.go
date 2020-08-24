package main

import (
	"fmt"
)

func main() {
	// 1
	x := 5
	y := 7
	sum := x + y
	// 2
	myslice := []int{5, 4, 3, 2, 1}
	myslice = append(myslice, 13)
	// 3
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Printf("x=%v,y=%v,sum=%v\n", x, y, sum)
	fmt.Printf("myslice=%v\n", myslice)
	fmt.Printf("vertices=%v\n", vertices)
	for key, value := range vertices {
		fmt.Printf("%v = %v\n", key, value)
	}
}
