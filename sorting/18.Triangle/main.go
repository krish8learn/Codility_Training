package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{10, 2, 5, 1, 8, 20}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	N := len(A)
	if N == 0 {
		return 0
	}

	sort.Ints(A)
	for i := 0; i < N-2; i++ {
		if A[i]+A[i+1] > A[i+2] && A[i]+A[i+2] > A[i] && A[i+2]+A[i+1] > A[i] {
			return 1
		}
	}
	return 0

}
