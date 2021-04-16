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
		//binary[i]--> rune(1-->49, 0-->48)
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
			//binary gap means 10001 -- gap is 3, if last element is not 1 after getting
			//consecutive 0, then it will not consider gap such as 10000--- no gap
			//so binary[j] -- last element and value is stil 0  we will not add to the slice
			//if not last element and loop has stopped because of occurence of 1 , then that
			//should be added to the slice
			//this is the tricky part
			if j != len(binary) {
				count = append(count, gap)
			}
			//if this condition is true means for loop has stopped because of occurence of 1
			//not because of last element

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
