package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10
	y := 85
	d := 30
	fmt.Println(solution(x, y, d))
}

func solution(x int, y int, d int) int {
	fvalue := float64(y-x) / float64(d)
	//fmt.Println(fvalue)
	jump := int(math.Ceil(fvalue))
	return jump
}
