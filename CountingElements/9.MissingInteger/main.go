package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 3, 6, 4, 1, 2}
	fmt.Println(Missinteger(A))
}

func Missinteger(A []int) int {
	list := removeDuplicte(A)
	fmt.Println(list)
	sort.Ints(list)
	fmt.Println(list)
	small := 1
	for contains(list, small) {
		small++
	}
	return small

}

func removeDuplicte(A []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range A {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func contains(list []int, small int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == small {
			return true
		}
	}
	return false
}
