package main

import "fmt"

func main() {
	A := []int{2, 3, 1, 5}
	fmt.Println(missing(A))
}

func missing(A []int) int {
	N := len(A) + 1
	//fmt.Println(N)
	ret := ((N + 1) * N) / 2
	for _,val := range A {
		fmt.Println(ret, val)
		ret -= val
	}
	return ret
}
