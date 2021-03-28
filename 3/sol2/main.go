package main

import "fmt"

func main() {
	A := []int{9, 9, 9, 9, 9}
	fmt.Println(unpaired(A))
}

func unpaired(A []int) int {
	ret := 0
	for i := 0; i < len(A); i++ {
		ret = ret ^ A[i]
	}
	return ret
}
