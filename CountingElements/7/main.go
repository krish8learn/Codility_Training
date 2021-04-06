package main

import "fmt"

func main() {
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	fmt.Println(frogRiver(5, A))
}

func frogRiver(X int, A []int) int {
	ret := -1
	var temp []int
	for i := 1; i <= X; i++ {
		temp = append(temp, i)
	}
	//fmt.Println(temp)
	//lengthA := len(A)

	for i := 0; i < len(A); i++ {
		//fmt.Println("comp: ", A[i])
		//removing the element
		for j := 0; j < len(temp); j++ {
			if temp[j] == A[i] {
				temp = remove(temp, j)
				//fmt.Println(temp)
				break
			}
		}
		//fmt.Println(i)
		if len(temp) == 0 {
			ret = i
			break
		}
	}
	//fmt.Println(ret)
	return ret
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
