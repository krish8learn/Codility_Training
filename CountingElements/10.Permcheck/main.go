package main

import "fmt"

func main() {
	A := []int{1,1}
	fmt.Println(PermCheck(A))
}

func PermCheck(A []int) int {
	ret := 0
	var perm []int
	for i := 1; i <= len(A); i++ {
		perm = append(perm, i)
	}
	//fmt.Println(perm)

	for _, value := range A {
		if !contains(value, perm) {
			return ret
		} else {
			perm = remove(value, perm)
			//fmt.Println(perm)
		}
	}
	//fmt.Println(len(perm))
	if len(perm) == 0 {
		return 1
	}
	return ret
}

func contains(value int, perm []int) bool {
	for _, val := range perm {
		if val == value {
			return true
		}
	}
	return false
}

func remove(value int, perm []int) []int {
	for index := 0; index < len(perm); index++ {
		if value == perm[index] {
			perm[len(perm)-1], perm[index] = perm[index], perm[len(perm)-1]
		}
		//fmt.Println(perm)
	}
	return perm[:len(perm)-1]
}
