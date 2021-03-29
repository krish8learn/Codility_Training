package main

import "fmt"

func main() {
	x := 10
	y := 85
	d := 30
	fmt.Println(solution(x, y, d))
}

func solution(x int, y int, d int) int {
	count := 0

	for x < y {
		x += d
		count++
	}
	return count
}
