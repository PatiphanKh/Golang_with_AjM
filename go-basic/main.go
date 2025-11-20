package main

import (
	"fmt"
	"go-basic/cal"
)

func main() {
	// s := []int{1, 2, 3, 4}
	sum, avg := cal.Add([]int{10,20,30,40}...)
	fmt.Printf("%v %v",sum, avg)
}