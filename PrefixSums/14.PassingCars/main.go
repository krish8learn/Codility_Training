package main

import "fmt"

func main() {
	A := []int{0, 1, 0, 1, 1}
	fmt.Println(PassingCars(A))
}
func PassingCars(A []int) int {
	zeros := 0
	ones := 0

	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			zeros++
		} else if A[i] == 1 {
			ones += zeros
		}
	}
	if ones > 1000000000 {
		return -1
	}
	return ones
}
