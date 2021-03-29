package main

import (
	"fmt"
	"math"
)

func main() {
	A := []int{3, 1, 2, 4, 3}
	fmt.Println(tapeequilibrium(A))
}

func tapeequilibrium(A []int) int {
	firstpart := A[0]
	//fmt.Println("first", firstpart)
	lastpart := 0
	for i := 1; i < len(A); i++ {
		lastpart += A[i]
	}
	//fmt.Println("last:", lastpart)
	diff := math.Abs(float64(firstpart) - float64(lastpart))
	//fmt.Println("diff",diff)
	for i := 1; i < len(A)-1; i++ {
		firstpart += A[i]
		lastpart -= A[i]
		temp := math.Abs(float64(firstpart) - float64(lastpart))
		if temp < diff {
			diff = temp
		}
	}
	return int(diff)
}
