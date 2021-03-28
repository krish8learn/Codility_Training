package main

import "fmt"

func main() {
	A := []int{9, 9, 9, 9, 9}
	fmt.Println(unpaired(A))
}

func unpaired(A []int) int {
	comp := 0
	//fmt.Println(A)
	if len(A) > 0 {
		for i := 0; i < len(A); i++ {
			comp = A[i]
			//fmt.Println(comp)
			count := 0
			for j := 0; j < len(A); j++ {
				if comp == A[j] {
					count++
					//fmt.Println(count)
				}
			}
			if count%2 != 0 {
				break
			} else{
				count = 0
				comp = 0
			}
		}
	}
	//fmt.Println("comp",comp)
	return comp
}
