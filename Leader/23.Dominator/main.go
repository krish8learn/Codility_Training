package main

import (
	"fmt"
	"sort"
)

func main() {
	//A := []int{3, 4, 3, 2, 3, -1, 3, 3}
	//A := []int{2,3,1}
	//A := []int{2, 2, 2, 2, 2}
	//A := []int{78, 2, 2, 2, 2, 2}
	//A := []int{}
	//A := []int{56, 12, 34, 21, -90, -90, 23, 54, 12, 58}
	//A := []int{1289, 786, 1289, 1892, 1289, 775, 1289, 1289, 1289}
	A := []int{3, 2, 2, 3}
	fmt.Println("Dominator Index: ", solution(A))

}

func solution(A []int) int {
	var duplicate = make([]int, len(A))
	length := copy(duplicate, A)
	if length == 0 {
		return -1
	} else if length == 1 {
		return 0
	}
	//fmt.Println("Length:", length)
	sort.Ints(duplicate)
	//fmt.Println("Sorted New Array:", duplicate)
	//fmt.Println("Old Array: ", A)
	half := length / 2
	//fmt.Println(half)
	count := 1
	element := 0
	for i := 1; i < length; i++ {
		if duplicate[i] == duplicate[i-1] {
			count++
			if count > half {
				element = duplicate[i]
				break
			}
		}
	}
	//fmt.Println("Element:", element)

	for i := 0; i < length; i++ {
		//fmt.Println(A[i])
		if element == A[i] {
			return i
		}
	}
	return -1
}
