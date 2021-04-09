package main

import (
	"fmt"
)

func main() {
	A := []int{4, 2, 2, 5, 1, 5, 8}
	//A := []int{10, 10, -1, 2, 4, -1, 2, -1}
	fmt.Println("RES ", MinAvg2Silice(A))
}

func MinAvg2Silice(A []int) int {
	N := len(A)
	minv := float64((A[0] + A[1]) / 2.0)
	result := 0

	for i := 0; i < N-2; i++ {
		tempmin := min(float64(((A[i] + A[i+1]) / 2.0)), float64(((A[i] + A[i+1] + A[i+2]) / 2.0)))

		if tempmin < minv {
			minv = tempmin
			result = i
		}
	}

	if float64((A[N-2]+A[N-1])/2.0) < minv {
		result = N - 2
	}

	return result
}

func min(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
