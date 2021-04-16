package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var A []int
	var K int
	var text string
	fmt.Println("Enter the elements of slice")
	fmt.Scanf("%s", &text)
	for _, val := range strings.Split(text, ",") {
		intval, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Not Integer value")
			break
		}
		A = append(A, intval)
	}
	fmt.Println("Enter how many times you want to rotate the array")
	fmt.Scanf("%d", &K)
	fmt.Printf("After %d times: %v \n", K, Rotate(A, K))
	//fmt.Println(Rotate(A, K))

}

func Rotate(A []int, K int) []int {
	for K != 0 {
		last := A[len(A)-1]
		for j := len(A) - 1; j > 0; j-- {
			A[j] = A[j-1]
		}
		A[0] = last
		K--
	}
	//fmt.Println(A)
	return A
}
