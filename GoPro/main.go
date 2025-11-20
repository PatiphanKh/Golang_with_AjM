package main

import (
	"fmt"
	"gopro/cal"
)

func main() {
	s := []int{1, 2, 3, 4}
	sum, avg := cal.Add(s)
	fmt.Printf("sum = %d", sum, "avg = %0.2f", avg)
}
