package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BinaryGap(32))
}

func BinaryGap(n int) int {
	//converting the integer into binary
	binary := strconv.FormatInt(int64(n), 2)
	//fmt.Println(binary)
	var count []int
	for i := 0; i < len(binary); i++ {
		//fmt.Println(binary[i])
		if binary[i] == 49 {
			//fmt.Println(" if true")
			gap := 0
			j := i + 1
			if j < len(binary) {
				for binary[j] != 49 {
					//fmt.Println("for true")
					j++
					if j == len(binary) {
						break
					}
					gap++
				}
			}
			//fmt.Println("gap", gap)
			if j != len(binary) {
				count = append(count, gap)
			}

		}
	}

	//fmt.Println(count)

	max := 0
	for _, value := range count {
		if value > max {
			max = value
		}
	}
	return max
}
