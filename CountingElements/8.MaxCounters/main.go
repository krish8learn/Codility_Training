package main

import "fmt"

func main() {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	N := 5
	fmt.Println(MaxCounter(N, A))
}

func MaxCounter(N int, A []int) []int {
	//setting up the counter with initial value 0
	var counter []int
	setmax := 0
	for i := 1; i <= N; i++ {
		counter = append(counter, 0)
	}
	for i := 0; i < len(A); i++ {
		if A[i] <= N && A[i] > 0 {
			counter[A[i]-1]++
			if setmax < counter[A[i]-1] {
				setmax = counter[A[i]-1]
			}
		} else if A[i] > N {
			for i := 0; i < N; i++ {
				counter[i] = setmax
			}
		}
	}
	return counter
}

func MaxValue(counter []int) int {
	max := 0
	for i := 0; i < len(counter); i++ {
		if counter[i] > max {
			max = counter[i]
		}
	}
	return max
}
