package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{2, 1, 1, 2, 3, 1}
	//A := []int{-3, -2, -1, 0, 1, 2, 4}
	//A := []int{10000234, 10000235, 10000236, 10000239}
	fmt.Println(Distinct(A))
}
func Distinct(A []int) int {
	N := len(A)
	if N == 0 {
		return 0
	}
	sort.Ints(A)
	count := 0
	for i := 0; i < N-1; i++ {
		//fmt.Println("Passing", i)
		if A[i] != A[i+1] {
			count++
		}
	}
	return count
}
