package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{-3, 1, 2, -2, 5, 6}

	fmt.Println(MaxProductThree(A))
}

func MaxProductThree(A []int) int {
	sort.Ints(A)
	N := len(A)
	max1 := A[N-1] * A[N-2] * A[N-3]
	max2 := A[0] * A[1] * A[N-1]

	if max2 > max1 {
		return max2
	}

	return max1

}
