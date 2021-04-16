package main

import (
	"fmt"
)

func main() {
	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}
	fmt.Println(GenomicRangeQuery(S, P, Q))
}

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	code := []string{"A","C","G","T"}
	
	var ans []int
	for i := range P {
		condition := false
		for index, value := range code{
			for _, x := range S[P[i]:Q[i]+1]{
				if value == string(x){
					condition = true
					ans = append(ans, index+1)
					break
				}
			}
			if condition == true{
				break
			}
		}
	}
	return ans
}
